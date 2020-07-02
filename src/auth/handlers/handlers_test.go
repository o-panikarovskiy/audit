package handlers

import (
	"audit/src/config"
	"audit/src/di"
	"audit/src/middlewares"
	"audit/src/sessions/mem"
	"audit/src/user"
	"audit/src/user/emailconfirm"
	"audit/src/user/testrep"
	"audit/src/user/userservice"
	"audit/src/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func init() {
	cfg := &config.AppConfig{
		SessionAgeSec: 60 * 60,
		Cookie: config.AppConfigCookie{
			Name:  "sid",
			Hash:  "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Block: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		},
	}
	rep := testrep.NewTestRepository()
	ms := mem.NewStorage()
	ec := emailconfirm.NewEmailConfirmService(cfg)
	srv := userservice.NewDefaultUserService(rep, ms, ec, cfg)

	deps := &di.ServiceLocator{}
	deps.Register(cfg)
	deps.Register(ms)
	deps.Register(srv)

	di.Set(deps)
}

func TestCheckSession(t *testing.T) {
	handler := http.HandlerFunc(CheckSession)

	t.Run("should be 401", func(t *testing.T) {
		rr, req := setupRecorder("GET", "/api/auth/check", nil)
		handler.ServeHTTP(rr, req)
		checkJSONResponse(t, rr, http.StatusUnauthorized)
	})

	t.Run("should be 401", func(t *testing.T) {
		rr, req := setupRecorder("GET", "/api/auth/check", nil)
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithSessionID("wrong sid")))
		checkJSONResponse(t, rr, http.StatusUnauthorized)
	})

	t.Run("should be 200", func(t *testing.T) {
		u, sid := addUser(t, "hadnler-auth-check-session@g.com", "12345678")
		rr, req := setupRecorder("GET", "/api/auth/check", nil)

		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithSessionID(sid)))

		checkJSONResponse(t, rr, http.StatusOK)
		checkCookieExists(t, rr, "sid")
		answer := parseBodyToJSON(t, rr)

		if answer["email"] != u.Email {
			t.Errorf("Wrong user email: got %v; want %v", answer["email"], u.Email)
		}

		if answer["id"] != u.ID {
			t.Errorf("Wrong user id: got %v; want %v", answer["id"], u.ID)
		}
	})
}

func TestSignIn(t *testing.T) {
	handler := http.HandlerFunc(SignIn)

	t.Run("should be 400", func(t *testing.T) {
		rr, req := setupRecorder("POST", "/api/auth/signin", nil)
		handler.ServeHTTP(rr, req)
		checkJSONResponse(t, rr, http.StatusBadRequest)
	})

	t.Run("should be 400", func(t *testing.T) {
		data := map[string]interface{}{"username": "", "password": ""}
		rr, req := setupRecorder("POST", "/api/auth/signin", data)
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithJSON(&data)))
		checkJSONResponse(t, rr, http.StatusBadRequest)
	})

	t.Run("should be 200", func(t *testing.T) {
		email := "hadnler-auth-signin@g.com"
		password := "12345678"
		u, sid := addUser(t, email, password)

		data := map[string]interface{}{"username": email, "password": password}
		rr, req := setupRecorder("POST", "/api/auth/signin", data)
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithSessionID(sid).WithJSON(&data)))

		checkJSONResponse(t, rr, http.StatusOK)
		checkCookieExists(t, rr, "sid")
		answer := parseBodyToJSON(t, rr)

		if answer["email"] != u.Email {
			t.Errorf("Wrong user email: got %v; want %v", answer["email"], u.Email)
		}

		if answer["id"] != u.ID {
			t.Errorf("Wrong user id: got %v; want %v", answer["id"], u.ID)
		}

	})
}

func TestSignUp(t *testing.T) {
	handler := http.HandlerFunc(SignUp)

	t.Run("should be 400", func(t *testing.T) {
		rr, req := setupRecorder("POST", "/api/auth/signup", nil)
		handler.ServeHTTP(rr, req)
		checkJSONResponse(t, rr, http.StatusBadRequest)
	})

	t.Run("should be 400", func(t *testing.T) {
		data := map[string]interface{}{"email": "", "password": ""}
		rr, req := setupRecorder("POST", "/api/auth/signup", data)
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithJSON(&data)))
		checkJSONResponse(t, rr, http.StatusBadRequest)
	})

	t.Run("should be 400", func(t *testing.T) {
		data := map[string]interface{}{"email": "wrong", "password": ""}
		rr, req := setupRecorder("POST", "/api/auth/signup", data)
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithJSON(&data)))
		checkJSONResponse(t, rr, http.StatusBadRequest)
	})

	t.Run("should be 400", func(t *testing.T) {
		data := map[string]interface{}{"email": "test@g.com", "password": "small"}
		rr, req := setupRecorder("POST", "/api/auth/signup", data)
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithJSON(&data)))
		checkJSONResponse(t, rr, http.StatusBadRequest)
	})

	t.Run("should be 204", func(t *testing.T) {
		email := "hadnler-auth-signup@g.com"
		password := "12345678"

		data := map[string]interface{}{"email": email, "password": password}
		rr, req := setupRecorder("POST", "/api/auth/signup", data)
		handler.ServeHTTP(rr, req.WithContext(middlewares.NewContext(req.Context()).WithJSON(&data)))

		checkJSONResponse(t, rr, http.StatusNoContent)
	})
}

func TestConfirm(t *testing.T) {
	handler := http.HandlerFunc(EndRegistration)

	t.Run("should be 400", func(t *testing.T) {
		rr, req := setupRecorder("GET", "/api/auth/confirm/null", nil)
		handler.ServeHTTP(rr, req)
		checkJSONResponse(t, rr, http.StatusBadRequest)
	})

	t.Run("should be 200", func(t *testing.T) {
		email := "hadnler-signup-confirm@g.com"
		password := "12345678"

		srv := di.GetUserService()
		tokenID, _, err := srv.SignUp(email, password)
		if err != nil {
			t.Fatal(err)
		}

		rr, req := setupRecorder("GET", "/api/auth/confirm/", map[string]string{"token": tokenID})
		handler.ServeHTTP(rr, req)

		t.Log(rr.Body.String(), tokenID)

		checkJSONResponse(t, rr, http.StatusOK)
		checkCookieExists(t, rr, "sid")
		answer := parseBodyToJSON(t, rr)

		if answer["email"] != email {
			t.Errorf("Wrong user email: got %v; want %v", answer["email"], email)
		}
	})
}

func setupRecorder(method string, url string, data interface{}) (*httptest.ResponseRecorder, *http.Request) {
	var req *http.Request

	if method == http.MethodPost || method == http.MethodPut {
		req = setupBodyRec(method, url, data)
	} else {
		if d, ok := data.(map[string]string); ok {
			req = setupNoBodyRec(method, url, d)
		} else {
			req = setupNoBodyRec(method, url, nil)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	ctx := middlewares.NewContext(req.Context())
	req = req.WithContext(ctx)

	return rr, req
}

func setupBodyRec(method string, url string, data interface{}) *http.Request {
	var dataReader io.Reader
	if data != nil {
		body, _ := json.Marshal(data)
		dataReader = bytes.NewReader(body)
	}
	req, _ := http.NewRequest(method, url, dataReader)
	return req
}

func setupNoBodyRec(method string, url string, data map[string]string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	if data != nil {
		req = mux.SetURLVars(req, data)
	}
	return req
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

func checkJSONResponse(t *testing.T, rr *httptest.ResponseRecorder, status int) {
	if rr.Code != status {
		t.Errorf("Wrong status code: got %v; want %v", rr.Code, status)
	}
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json; charset=utf-8" {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
}

func checkCookieExists(t *testing.T, rr *httptest.ResponseRecorder, name string) {
	header := rr.Header().Get("Set-Cookie")
	cookies := strings.Split(header, ";")
	prefix := name + "="

	var found bool
	for _, v := range cookies {
		if strings.HasPrefix(v, prefix) {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Cookie %s does not exist in response headers", name)
	}
}

func parseBodyToJSON(t *testing.T, rr *httptest.ResponseRecorder) map[string]interface{} {
	rp, err := utils.JSONParse(rr.Body.String())
	if err != nil {
		t.Fatal(err)
	}

	return *rp
}
