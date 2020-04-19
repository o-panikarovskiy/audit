package user

import (
	"log"
)

type testRepo struct {
	users []*User
}

//NewTestRepository create new repository
func NewTestRepository() IRepository {
	log.Println("Create users table...")

	rep := &testRepo{
		users: make([]*User, 0),
	}

	rep.Store(&User{
		Email: "oleg.pnk@gmail.com",
		Role:  UserRole,
	})

	return rep
}

func (r *testRepo) ShutDown() {
	log.Println("Close users db connections...")
}

func (r *testRepo) Find(id string) (*User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

func (r *testRepo) FindByEmail(email string) (*User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, nil
}

func (r *testRepo) FindAll() ([]*User, error) {
	return r.users, nil
}

func (r *testRepo) Store(user *User) (*User, error) {
	r.users = append(r.users, user)
	return user, nil
}

func (r *testRepo) Update(user *User) (*User, error) {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return u, nil
		}
	}

	return nil, nil
}
