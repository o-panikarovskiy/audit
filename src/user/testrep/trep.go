package testrep

import (
	"audit/src/user"
	"audit/src/utils"
	"log"
	"time"
)

type testRepo struct {
	users []*user.User
}

//NewTestRepository create new repository
func NewTestRepository() user.IRepository {
	log.Println("Create users table...")

	rep := &testRepo{
		users: make([]*user.User, 0),
	}

	return rep
}

func (r *testRepo) FindByID(id string) (*user.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

func (r *testRepo) FindByUsername(email string) (*user.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, nil
}

func (r *testRepo) FindAll() ([]*user.User, error) {
	return r.users, nil
}

func (r *testRepo) Store(user *user.User) (*user.User, error) {
	user.ID = utils.CreateGUID()
	user.Created = time.Now()
	r.users = append(r.users, user)
	return user, nil
}

func (r *testRepo) Update(user *user.User) (*user.User, error) {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return u, nil
		}
	}

	return nil, nil
}
