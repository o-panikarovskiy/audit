package userservice

import (
	"audit/src/user"
	"audit/src/utils"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type signUpUserData struct {
	Email        string
	Role         int
	PasswordSalt string
	PasswordHash string
}

func (s *userService) SignUp(email string, password string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", invalidReqModelErr
	}

	email = strings.ToLower(email)
	exUser, err := s.FindByUsername(email)
	if err != nil {
		return "", "", err
	}

	if exUser != nil {
		return "", "", userExistsError
	}

	tokenID, err := s.storeSignUpData(email, password)
	if err != nil {
		return "", "", err
	}

	err = s.confirmator.Send(&user.User{Email: email}, tokenID, "")
	if err != nil {
		return "", "", err
	}

	return tokenID, "", nil
}

func (s *userService) EndSignUp(tokenID string, tokenValue string) (*user.User, string, error) {
	if tokenID == "" {
		return nil, "", badTokenError
	}

	usr, err := s.restoreSignUpData(tokenID)
	if err != nil {
		return nil, "", err
	}

	dbUser, err := s.Store(usr)
	if err != nil {
		return nil, "", err
	}

	_, err = s.sessions.Delete(confirmEmailKey + tokenID)
	if err != nil {
		return nil, "", err
	}

	sid := utils.RandomString(64)
	err = s.saveAuthSession(sid, dbUser)
	if err != nil {
		return nil, "", err
	}

	err = s.destroySignUpData(tokenID)
	if err != nil {
		return nil, "", err
	}

	return dbUser, sid, nil
}

func (s *userService) storeSignUpData(email string, password string) (string, error) {
	expiration := 24 * 60 * 60
	salt := utils.RandomString(64)
	tokenID := utils.RandomString(32)

	data := &signUpUserData{
		Email:        email,
		Role:         user.UserRole,
		PasswordSalt: salt,
		PasswordHash: utils.SHA512(password, salt),
	}

	err := s.sessions.SetJSON(confirmEmailKey+tokenID, data, expiration)
	if err != nil {
		return "", err
	}

	return tokenID, nil
}

func (s *userService) restoreSignUpData(tokenID string) (*user.User, error) {
	json, err := s.sessions.GetJSON(confirmEmailKey + tokenID)
	if err != nil {
		return nil, err
	}

	if json == nil {
		return nil, badTokenError
	}

	var data signUpUserData
	err = mapstructure.Decode(json, &data)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Email:        data.Email,
		Status:       "confirmed",
		PasswordHash: data.PasswordHash,
		PasswordSalt: data.PasswordSalt,
		Role:         data.Role,
	}

	return u, nil
}

func (s *userService) destroySignUpData(tokenID string) error {
	_, err := s.sessions.Delete(confirmEmailKey + tokenID)
	if err != nil {
		return err
	}
	return nil
}
