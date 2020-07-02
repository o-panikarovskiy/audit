package pgrep

import (
	"audit/src/user"

	"database/sql"
)

type pgRepository struct {
	db *sql.DB
}

//NewRepository create new pgRepository
func NewRepository(db *sql.DB) (user.IRepository, error) {
	rep := &pgRepository{db: db}

	err := rep.createUsersTable()
	if err != nil {
		return nil, err
	}

	return rep, nil
}
