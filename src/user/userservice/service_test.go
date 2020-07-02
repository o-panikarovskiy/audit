package userservice

import (
	"audit/src/config"
	"audit/src/sessions/mem"
	"audit/src/user"
	"audit/src/user/emailconfirmator"
	"audit/src/user/testrep"
	"audit/src/utils"
	"testing"
)

var cfg = &config.AppConfig{SessionAgeSec: 60 * 60}
var rep = testrep.NewTestRepository()
var ms = mem.NewStorage()
var ec = emailconfirmator.NewEmailConfirmatorService(cfg)
var srv = NewDefaultUserService(rep, ms, ec, cfg)

func TestSignUp(t *testing.T) {
	email := "signup@user.com"
	password := "12345678"

	t.Run("empty username", func(t *testing.T) {
		_, _, err := srv.SignUp("", "")
		if err == nil {
			t.Errorf(`SignUp("", "") hasn't error; want error`)
		} else if e, ok := err.(*utils.AppError); !ok {
			t.Errorf(`SignUp("", "") hasn wrong error %w; want AppError`, err)
		} else if e.Code != "INVALID_REQUEST_MODEL" {
			t.Errorf(`SignUp("", "") hasn wrong error code %s; want "INVALID_REQUEST_MODEL"`, e.Code)
		}
	})

	t.Run("good sign up", func(t *testing.T) {
		u, sid, tid, tval := signUp(t, email, password)

		_, _, err := srv.EndSignUp(tid, tval)
		if err == nil {
			t.Errorf(`Double run EndSignUp("%s", "%s") hasn't error; want error`, tid, tval)
		} else if _, ok := err.(*utils.AppError); !ok {
			t.Errorf(`Double run EndSignUp("%s", "%s") hasn unknown error %w; want AppError`, tid, tval, err)
		}

		au, _ := auth(t, email, password)
		if u.Email != au.Email {
			t.Errorf(`got auth user email %s; want %s`, au.Email, u.Email)
		}

		cu := checkSession(t, sid)
		if u.Email != cu.Email {
			t.Errorf(`got check session user email %s; want %s`, cu.Email, u.Email)
		}
	})

	t.Run("user already exists", func(t *testing.T) {
		_, _, err := srv.SignUp(email, password)
		if err == nil {
			t.Errorf(`SignUp("%s", "%s") hasn't error; want error`, email, password)
		} else if e, ok := err.(*utils.AppError); !ok {
			t.Errorf(`SignUp("%s", "%s") hasn unknown error %w; want AppError`, email, password, err)
		} else if e.Code != "USER_EXISTS" {
			t.Errorf(`SignUp("%s", "%s") hasn unknown error code %s; want "USER_EXISTS"`, email, password, e.Code)
		}
	})
}

func TestAuth(t *testing.T) {
	email := "auth@user.com"
	password := "12345678"

	t.Run("empty username", func(t *testing.T) {
		_, _, err := srv.Auth("", "")
		if err == nil {
			t.Errorf(`Auth("", "") hasn't error; want error`)
		} else if e, ok := err.(*utils.AppError); !ok {
			t.Errorf(`Auth("", "") hasn wrong error %w; want AppError`, err)
		} else if e.Code != "INVALID_REQUEST_MODEL" {
			t.Errorf(`Auth("", "") hasn wrong error code %s; want "INVALID_REQUEST_MODEL"`, e.Code)
		}
	})

	t.Run("good auth", func(t *testing.T) {
		u, _, _, _ := signUp(t, email, password)

		au, sid := auth(t, email, password)
		if u.Email != au.Email {
			t.Errorf(`got auth user email %s; want %s`, au.Email, u.Email)
		}

		cu := checkSession(t, sid)

		if u.Email != cu.Email {
			t.Errorf(`got check session user email %s; want %s`, cu.Email, u.Email)
		}
	})
}

func TestCheckAuthSession(t *testing.T) {
	email := "check-session@user.com"
	password := "12345678"

	t.Run("empty sid", func(t *testing.T) {
		_, err := srv.CheckAuthSession("")
		if err == nil {
			t.Errorf(`CheckAuthSession("") hasn't error; want error`)
		} else if e, ok := err.(*utils.AppError); !ok {
			t.Errorf(`CheckAuthSession("") hasn wrong error %w; want AppError`, err)
		} else if e.Code != "AUTH_ERROR" {
			t.Errorf(`CheckAuthSession("") hasn wrong error code %s; want "AUTH_ERROR"`, e.Code)
		}
	})

	t.Run("wrong sid", func(t *testing.T) {
		_, err := srv.CheckAuthSession("wrong sid")
		if err == nil {
			t.Errorf(`CheckAuthSession("wrong sid") hasn't error; want error`)
		} else if e, ok := err.(*utils.AppError); !ok {
			t.Errorf(`CheckAuthSession("wrong sid") hasn wrong error %w; want AppError`, err)
		} else if e.Code != "AUTH_ERROR" {
			t.Errorf(`CheckAuthSession("wrong sid") hasn wrong error code %s; want "AUTH_ERROR"`, e.Code)
		}
	})

	t.Run("good sid", func(t *testing.T) {
		u, sid, _, _ := signUp(t, email, password)
		cu := checkSession(t, sid)

		if u.Email != cu.Email {
			t.Errorf(`Wrong users emails beetween EndSignUp and CheckAuthSession: %s and %s`, u.Email, cu.Email)
		}
	})
}

func TestRestoreSessionUser(t *testing.T) {
	email := "restore-session-user@user.com"
	password := "12345678"

	t.Run("empty sid", func(t *testing.T) {
		u, err := srv.RestoreSessionUser("")
		if err == nil && u != nil {
			t.Errorf(`RestoreSessionUser("") has empty error and non empy user; want empty user`)
		}
	})

	t.Run("wrong sid", func(t *testing.T) {
		u, err := srv.RestoreSessionUser("wrong sid")
		if err == nil && u != nil {
			t.Errorf(`RestoreSessionUser("wrong sid") has empty error and non empy user; want empty user`)
		}
	})

	t.Run("good sid", func(t *testing.T) {
		_, sid, _, _ := signUp(t, email, password)
		u, err := srv.RestoreSessionUser(sid)

		if err != nil {
			t.Errorf(`RestoreSessionUser("%s") has error; want no error`, sid)
		}

		if u.Email != email {
			t.Errorf(`Wrong users emails beetween EndSignUp and RestoreSessionUser: %s and %s`, u.Email, email)
		}
	})
}

func TestGetSessionStorage(t *testing.T) {
	t.Run("must be equal", func(t *testing.T) {
		if ms != srv.GetSessionStorage() {
			t.Fail()
		}
	})
}

func signUp(t *testing.T, email string, password string) (*user.User, string, string, string) {
	tid, tval, err := srv.SignUp(email, password)
	if err != nil {
		t.Errorf(`SignUp("%s", "%s") has error %w; want no error`, email, password, err)
	}

	u, sid, err := srv.EndSignUp(tid, tval)
	if err != nil {
		t.Errorf(`EndSignUp("%s", "%s") has error %w; want no error`, tid, tval, err)
	}

	if sid == "" {
		t.Errorf(`EndSignUp("%s", "%s") sid is empty; want no empty`, tid, tval)
	}

	if u == nil {
		t.Errorf(`EndSignUp("%s", "%s") user is empty; want no empty`, tid, tval)
	}

	if u.Email != email {
		t.Errorf(`got user email %s; want %s`, u.Email, email)
	}

	return u, sid, tid, tval
}

func auth(t *testing.T, email string, password string) (*user.User, string) {
	au, sid, err := srv.Auth(email, password)
	if err != nil {
		t.Errorf(`Auth("%s", "%s") has error %w; should work correctly`, email, password, err)
	}
	if au == nil {
		t.Errorf(`Auth("%s", "%s") user is empty; want no empty`, email, password)
	}

	return au, sid
}

func checkSession(t *testing.T, sid string) *user.User {
	cu, err := srv.CheckAuthSession(sid)
	if e, ok := err.(*utils.AppError); err != nil && !ok {
		t.Errorf(`CheckAuthSession("%s") has wrong error %w; want AppError`, sid, err)
	} else if err != nil && e.Code != "AUTH_ERROR" {
		t.Errorf(`CheckAuthSession("%s") has wrong error code %s; want "AUTH_ERROR"`, sid, e.Code)
	}
	if cu == nil {
		t.Errorf(`CheckAuthSession("%s") user is empty; want non empty`, sid)
	}
	return cu
}
