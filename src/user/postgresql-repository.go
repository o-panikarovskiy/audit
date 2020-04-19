package user

import (
	"log"
)

type postgreSQLrepo struct {
	pool interface{}
}

//NewPostgreSQLRepository create new repository
func NewPostgreSQLRepository(pool interface{}) IRepository {
	rep := &postgreSQLrepo{pool: pool}
	log.Println("Create users table...")
	return rep
}

func (r *postgreSQLrepo) ShutDown() {
	log.Println("Close users db connections...")
}

func (r *postgreSQLrepo) Find(id string) (*User, error) {
	return nil, nil
}

func (r *postgreSQLrepo) FindByEmail(email string) (*User, error) {
	return nil, nil
}

func (r *postgreSQLrepo) FindAll() ([]*User, error) {
	return nil, nil
}

func (r *postgreSQLrepo) Store(user *User) (*User, error) {
	return nil, nil
}

func (r *postgreSQLrepo) Update(user *User) (*User, error) {
	return nil, nil
}
