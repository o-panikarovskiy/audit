package userservice

import (
	"audit/src/config"
	"audit/src/sessions"
	"audit/src/user"
	"audit/src/utils"
)

const authSidKey = "AUTH:SID:"
const authUserKey = "AUTH:USER:"
const confirmEmailKey = "AUTH:EMAIL:"

var badTokenError = &utils.AppError{
	Code:    "BAD_TOKEN",
	Message: "Token not found",
}

var userExistsError = &utils.AppError{
	Code:    "USER_EXISTS",
	Message: "User already exists",
}

var invalidReqModelErr = &utils.AppError{
	Code:    "INVALID_REQUEST_MODEL",
	Message: "Email or password is incorrect",
}

var authAppErr = &utils.AppError{
	Code:    "AUTH_ERROR",
	Message: "Email or password is incorrect",
}

type userService struct {
	repo        user.IRepository
	sessions    sessions.IStorage
	confirmator user.IConfirmService
	cfg         *config.AppConfig
}

// NewDefaultUserService create new repository
func NewDefaultUserService(
	rep user.IRepository,
	ses sessions.IStorage,
	cfrm user.IConfirmService,
	cfg *config.AppConfig,
) user.IService {
	return &userService{
		repo:        rep,
		sessions:    ses,
		confirmator: cfrm,
		cfg:         cfg,
	}
}

func (s *userService) FindByID(id string) (*user.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) FindByUsername(email string) (*user.User, error) {
	return s.repo.FindByUsername(email)
}

func (s *userService) FindAll() ([]*user.User, error) {
	return s.repo.FindAll()
}

func (s *userService) Store(user *user.User) (*user.User, error) {
	return s.repo.Store(user)
}

func (s *userService) Update(user *user.User) (*user.User, error) {
	return s.repo.Update(user)
}

func (s *userService) GetSessionStorage() sessions.IStorage {
	return s.sessions
}
