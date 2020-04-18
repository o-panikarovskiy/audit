package user

import "audit/src/utils"

type defaultService struct {
	repo IRepository
}

//NewUserService create new repository
func NewUserService(r IRepository) IService {
	return &defaultService{repo: r}
}

func (s *defaultService) Find(id string) (*User, error) {
	return s.repo.Find(id)
}

func (s *defaultService) FindAll() ([]*User, error) {
	return s.repo.FindAll()
}

func (s *defaultService) Save(user *User) (*User, error) {
	return s.repo.Save(user)
}

func (s *defaultService) Register(user *User) (string, error) {
	return utils.CreateGUID(), nil
}

func (s *defaultService) Auth(user *User, password string) error {
	return nil
}

func (s *defaultService) CheckSession(sessionID string) (*User, error) {
	return s.Find(sessionID)
}

func (s *defaultService) GetRepo() IRepository {
	return s.repo
}
