package defservice

import (
	"audit/src/config"
	"audit/src/sessions/mem"
	"audit/src/user"
	"audit/src/user/emailconfirm"
	"audit/src/user/testrep"
	"audit/src/utils"
	"testing"
)

var cfg = &config.AppConfig{SessionAgeMin: 60}
var rep = testrep.NewTestRepository()
var ms = mem.NewStorage()
var ec = emailconfirm.NewEmailConfirmService(cfg)
var srv = NewDefaultUserService(rep, ms, ec, cfg)

type signUpTestCase struct {
	email    string
	password string
	appError string
}

var signUpCases = []signUpTestCase{
	{
		email:    "exists-user@t.com",
		password: "testpass",
		appError: "USER_EXISTS",
	},
	{
		email:    "test1@t.com",
		password: "testpass",
		appError: "",
	},
	{
		email:    "test2@t.com",
		password: "testpass",
		appError: "",
	},
	{
		email:    "test1@t.com",
		password: "testpass2",
		appError: "USER_EXISTS",
	},
}

func TestSignUp(t *testing.T) {
	rep.Store(&user.User{Email: signUpCases[0].email})

	for i, c := range signUpCases {
		tokenID, tokenVal, err := srv.SignUp(c.email, c.password)

		if c.appError == "" {
			if err != nil {
				t.Errorf("#%d: SignUp(%s, %s) has error %w; want no error", i, c.email, c.password, err)
			}
			if tokenID == "" {
				t.Errorf("tokenID is non empty; want empty")
			}
		} else {
			if err == nil {
				t.Errorf("#%d: SignUp(%s, %s) has no error; want %s", i, c.email, c.password, c.appError)
			} else if e, ok := err.(*utils.AppError); !ok {
				t.Errorf("#%d: SignUp(%s, %s) has error %w; want %s", i, c.email, c.password, err, c.appError)
			} else if e.Code != c.appError {
				t.Errorf("#%d: SignUp(%s, %s) has different error.Code: %s; want %s", i, c.email, c.password, e.Code, c.appError)
			}
			if tokenID != "" {
				t.Errorf("tokenID is non empty; want empty")
			}
		}

		if tokenID == "" {
			continue
		}

		usr, sid, err := srv.EndSignUp(tokenID, tokenVal)
		if err != nil {
			t.Errorf("#%d: EndSignUp(%s, %s) has error %w; want no error", i, tokenID, tokenVal, err)
		}
		if sid == "" {
			t.Errorf("#%d: EndSignUp(%s, %s) api sid is empty; want non empty", i, tokenID, tokenVal)
		}
		if usr == nil {
			t.Errorf("#%d: EndSignUp(%s, %s) user is empty; want non empty", i, tokenID, tokenVal)
		}
	}
}
