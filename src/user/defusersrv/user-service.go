package defusersrv

import (
	"audit/src/config"
	"audit/src/sessions"
	"audit/src/user"
)

type userService struct {
	repo     user.IRepository
	sessions sessions.IStorage
	cfg      *config.AppConfig
}

//NewDefaultUserService create new repository
func NewDefaultUserService(r user.IRepository, s sessions.IStorage, cfg *config.AppConfig) user.IService {
	return &userService{
		repo:     r,
		sessions: s,
		cfg:      cfg,
	}
}

func (s *userService) Find(id string) (*user.User, error) {
	return s.repo.Find(id)
}

func (s *userService) FindByEmail(email string) (*user.User, error) {
	return s.repo.FindByEmail(email)
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

func (s *userService) ShutDown() {
	s.repo.ShutDown()
}
