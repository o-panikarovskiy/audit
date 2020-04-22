package pgrep

import (
	"audit/src/user"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type pgRepository struct {
	pool *pgxpool.Pool
}

//NewRepository create new pgRepository
func NewRepository(db *pgxpool.Pool) (user.IRepository, error) {
	rep := &pgRepository{pool: db}

	log.Println("Create users table...")

	err := rep.createUsersTable()
	if err != nil {
		return nil, err
	}

	return rep, nil
}
