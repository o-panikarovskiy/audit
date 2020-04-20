package user

import (
	"audit/src/utils"
)

type userService struct {
	repo IRepository
}

//NewUserService create new repository
func NewUserService(r IRepository) IService {
	return &userService{repo: r}
}

func (s *userService) Auth(email string, password string) (*User, error) {
	user, err := s.FindByEmail(email)
	if err != nil {
		return nil, utils.NewAppError("APP_ERROR", err.Error())
	}

	if user == nil ||
		user.PasswordHash != utils.SHA512(password, user.PasswordSalt) {
		return nil, utils.NewAppError("AUTH_ERROR", "Email or password is incorrect")
	}

	return user, nil
}

func (s *userService) CheckSession(sessionID string) (*User, error) {
	return s.Find(sessionID)
}

func (s *userService) Register(email string, password string) (*User, error) {
	exUser, err := s.FindByEmail(email)
	if err != nil {
		return nil, utils.NewAppError("APP_ERROR", err.Error())
	}

	if exUser != nil {
		return nil, utils.NewAppError("USER_EXISTS", "User already exists")
	}

	salt := utils.RandomString(64)
	user := &User{
		ID:           utils.CreateGUID(),
		Email:        email,
		Role:         UserRole,
		PasswordSalt: salt,
		PasswordHash: utils.SHA512(password, salt),
	}

	s.Store(user)

	return user, nil
}

func (s *userService) Find(id string) (*User, error) {
	return s.repo.Find(id)
}

func (s *userService) FindByEmail(email string) (*User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) FindAll() ([]*User, error) {
	return s.repo.FindAll()
}

func (s *userService) Store(user *User) (*User, error) {
	return s.repo.Store(user)
}

func (s *userService) Update(user *User) (*User, error) {
	return s.repo.Update(user)
}

func (s *userService) GetRepo() IRepository {
	return s.repo
}
