package user

import "audit/src/utils"

type repo struct {
	pool interface{}
}

//NewPostgreSQLRepository create new repository
func NewPostgreSQLRepository(pool interface{}) IRepository {
	return &repo{pool: pool}
}

func (r *repo) Find(id string) (*User, error) {
	u := &User{
		ID:   utils.CreateGUID(),
		Role: AdminRole,
	}

	return u, nil
}

func (r *repo) FindAll() ([]*User, error) {
	res := make([]*User, 0)
	return res, nil
}

func (r *repo) Save(user *User) (*User, error) {
	return nil, nil
}

func (r *repo) ShutDown() {

}
