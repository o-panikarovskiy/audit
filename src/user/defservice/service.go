package defservice

import (
	"audit/src/config"
	"audit/src/sessions"
	"audit/src/user"
)

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

func (s *userService) GetRepo() user.IRepository {
	return s.repo
}

func (s *userService) GetSessionStorage() sessions.IStorage {
	return s.sessions
}
