package defservice

import (
	"audit/src/config"
	"audit/src/sessions/mem"
	"audit/src/user/emailconfirm"
	"audit/src/user/testrep"
	"testing"
)

func TestSignUp(t *testing.T) {
	cfg := &config.AppConfig{SessionAgeMin: 60}
	rep := testrep.NewTestRepository()
	ms := mem.NewStorage()
	ec := emailconfirm.NewEmailConfirmService(cfg)
	srv := NewDefaultUserService(rep, ms, ec, cfg)

	var confirmTokenID string
	var email string = "testuser@t.com"
	var password string = "testpass"

	t.Run("returns sid and user", func(t *testing.T) {
		tID, tval, err := srv.SignUp(email, password)
		if err != nil {
			t.Error(err)
			return
		}

		if tval != "" {
			t.Errorf("%v should be empty", tval)
			return
		}

		if tID == "" {
			t.Errorf("confirmTokenID should be not empty")
			return
		}

		confirmTokenID = tID
	})

	t.Run("should complete registration", func(t *testing.T) {
		if confirmTokenID == "" {
			t.Errorf("confirmTokenID should be not empty")
			return
		}

		usr, apiSID, err := srv.EndSignUp(confirmTokenID, "")
		if err != nil {
			t.Error(err)
			return
		}

		if apiSID == "" {
			t.Errorf("apiSID should be not empty")
			return
		}

		if usr == nil {
			t.Errorf("user should be not empty")
			return
		}

		if usr.Email != email {
			t.Errorf("got %q, want %q", usr.Email, email)
		}
	})
}
