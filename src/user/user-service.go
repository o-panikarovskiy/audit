package user

import (
	"audit/src/config"
	"audit/src/sessions"
	"audit/src/utils"
)

type userService struct {
	repo     IRepository
	sessions sessions.IStorage
	cfg      *config.AppConfig
}

//NewUserService create new repository
func NewUserService(r IRepository, s sessions.IStorage, cfg *config.AppConfig) IService {
	return &userService{
		repo:     r,
		sessions: s,
		cfg:      cfg,
	}
}

func (s *userService) Auth(email string, password string) (*User, string, error) {
	user, err := s.FindByEmail(email)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	if user == nil ||
		user.PasswordHash != utils.SHA512(password, user.PasswordSalt) {
		return nil, "", utils.NewAppError("AUTH_ERROR", "Email or password is incorrect")
	}

	_, err = s.destroyUserSession(user.ID)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	sid, err := s.saveUserSession(user.ID)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	return user, sid, nil
}

func (s *userService) CheckSession(sessionID string) (*User, error) {
	return s.Find(sessionID)
}

func (s *userService) Register(email string, password string) (*User, string, error) {
	exUser, err := s.FindByEmail(email)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	if exUser != nil {
		return nil, "", utils.NewAppError("USER_EXISTS", "User already exists")
	}

	salt := utils.RandomString(64)
	user := &User{
		ID:           utils.CreateGUID(),
		Email:        email,
		Role:         UserRole,
		PasswordSalt: salt,
		PasswordHash: utils.SHA512(password, salt),
	}

	sid, err := s.saveUserSession(user.ID)
	if err != nil {
		return nil, "", utils.NewAppError("APP_ERROR", err.Error())
	}

	s.Store(user)
	return user, sid, nil
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

func (s *userService) GetSessionStorage() sessions.IStorage {
	return s.sessions
}

func (s *userService) ShutDown() {
	s.repo.ShutDown()
}

func (s *userService) destroyUserSession(userID string) (string, error) {
	sid, err := s.sessions.Get(userID)

	if err != nil {
		return "", nil
	}

	if sid != "" {
		err = s.sessions.Delete(sid)
	}

	if err != nil {
		return "", nil
	}

	return sid, nil
}

func (s *userService) saveUserSession(userID string) (string, error) {
	sid := utils.RandomString(64)

	err := s.sessions.Set(sid, userID, s.cfg.SessionAge)
	if err != nil {
		return "", nil
	}

	err = s.sessions.Set(userID, sid, s.cfg.SessionAge)
	if err != nil {
		return "", nil
	}

	return sid, nil
}
