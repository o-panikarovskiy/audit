package defservice

import (
	"audit/src/user"
	"audit/src/utils"
	"strings"

	"github.com/mitchellh/mapstructure"
)

const confirmEmailKey = "AUTH:REG:"

type signUpUserData struct {
	Email        string
	Role         int
	PasswordSalt string
	PasswordHash string
}

func (s *userService) SignUp(email string, password string) (string, string, error) {
	email = strings.ToLower(email)
	err := s.checkUserNotExists(email)
	if err != nil {
		return "", "", err
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
	usr, err := s.restoreSignUpUser(tokenID)
	if err != nil {
		return nil, "", err
	}

	dbUser, err := s.Store(usr)
	if err != nil {
		return nil, "", &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	_, err = s.sessions.Delete(confirmEmailKey + tokenID)
	if err != nil {
		return nil, "", &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	sid := utils.RandomString(64)
	err = s.saveAuthSession(sid, dbUser)
	if err != nil {
		return nil, "", &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	return dbUser, "", nil
}

func (s *userService) restoreSignUpUser(tokenID string) (*user.User, error) {
	json, err := s.sessions.GetJSON(confirmEmailKey + tokenID)
	if err != nil {
		return nil, utils.NewAppError("APP_ERROR", err.Error())
	}

	if json == nil {
		return nil, &utils.AppError{Code: "BAD_TOKEN_ID", Message: "Token not found"}
	}

	var data signUpUserData
	err = mapstructure.Decode(json, &data)
	if err != nil {
		return nil, &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
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
		return "", &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	return tokenID, nil
}

func (s *userService) checkUserNotExists(email string) error {
	exUser, err := s.FindByUsername(email)

	if err != nil {
		return &utils.AppError{Code: "APP_ERROR", Message: err.Error(), Err: err}
	}

	if exUser != nil {
		return &utils.AppError{Code: "USER_EXISTS", Message: "User already exists"}
	}

	return nil
}
