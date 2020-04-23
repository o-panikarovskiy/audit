package handlers

import (
	"audit/src/config"
	"audit/src/di"
	"audit/src/middlewares"
	"audit/src/sessions/mem"
	"audit/src/user"
	"audit/src/user/defservice"
	"audit/src/user/emailconfirm"
	"audit/src/user/testrep"
	"audit/src/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	cfg := &config.AppConfig{SessionAgeMin: 60}
	rep := testrep.NewTestRepository()
	ms := mem.NewStorage()
	ec := emailconfirm.NewEmailConfirmService(cfg)
	srv := defservice.NewDefaultUserService(rep, ms, ec, cfg)

	deps := &di.ServiceLocator{}
	deps.Register(cfg)
	deps.Register(ms)
	deps.Register(srv)

	di.Set(deps)
}

func setupRecorder(t *testing.T, method string, url string) (*httptest.ResponseRecorder, *http.Request) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ctx := middlewares.NewContext(req.Context())
	req = req.WithContext(ctx)

	return rr, req
}

func addUser(t *testing.T, email string, password string) (*user.User, string) {
	srv := di.GetUserService()
	tid, tval, err := srv.SignUp(email, password)
	if err != nil {
		t.Fatal(err)
	}

	u, sid, err := srv.EndSignUp(tid, tval)
	if err != nil {
		t.Fatal(err)
	}
	return u, sid
}

func TestCheckSession(t *testing.T) {
	handler := http.HandlerFunc(CheckSession)

	t.Run("should be 401", func(t *testing.T) {
		rr, req := setupRecorder(t, "GET", "/api/auth/check")
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Wrong status code: got %v; want %v", rr.Code, http.StatusUnauthorized)
		}
	})

	t.Run("should be 401", func(t *testing.T) {
		rr, req := setupRecorder(t, "GET", "/api/auth/check")
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithSessionID("wrong sid")))
		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Wrong status code: got %v; want %v", rr.Code, http.StatusUnauthorized)
		}
	})

	t.Run("should be 200", func(t *testing.T) {
		u, sid := addUser(t, "hadnler-check-session@g.com", "12345678")

		rr, req := setupRecorder(t, "GET", "/api/auth/check")
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithSessionID(sid)))
		if rr.Code != http.StatusOK {
			t.Errorf("Wrong status code: got %v; want %v", rr.Code, http.StatusOK)
		}

		rp, err := utils.JSONParse(rr.Body.String())
		if err != nil {
			t.Fatal(err)
		}

		ru := *rp
		if ru["email"] != u.Email {
			t.Errorf("Wrong user email: got %v; want %v", ru["email"], u.Email)
		}

		if ru["id"] != u.ID {
			t.Errorf("Wrong user id: got %v; want %v", ru["id"], u.ID)
		}
	})
}
