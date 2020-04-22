package pgrep

import (
	"audit/src/user"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type pgRep struct {
	pool *pgxpool.Pool
}

//NewRepository create new repository
func NewRepository(db *pgxpool.Pool) (user.IRepository, error) {
	rep := &pgRep{pool: db}

	log.Println("Create users table...")

	err := rep.createUsersTable()
	if err != nil {
		return nil, err
	}

	return rep, nil
}

func (r *pgRep) Find(id string) (*user.User, error) {
	return nil, nil
}

func (r *pgRep) FindByEmail(email string) (*user.User, error) {
	sql := r.getSelectQueryText() + ` WHERE "email" = $1 LIMIT 1`

	res, err := r.query(sql, email)
	if err != nil {
		return nil, err
	}

	log.Println(len(res))

	if len(res) > 0 {
		return res[0], nil
	}

	return nil, nil
}

func (r *pgRep) FindAll() ([]*user.User, error) {
	return nil, nil
}

func (r *pgRep) Store(u *user.User) (*user.User, error) {
	sql := `INSERT INTO "public"."users"(					                              
					                              "name", 					                             
					                              "email", 
					                              "status", 
					                              "passwordHash",
					                              "passwordSalt",
					                              "role" 
					) 
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING *;`

	res, err := r.query(sql, u.Name, u.Email, u.Status, u.PasswordHash, u.PasswordSalt, u.Role)

	if err != nil {
		return nil, err
	}

	return res[0], nil
}

func (r *pgRep) Update(u *user.User) (*user.User, error) {
	return nil, nil
}

func (r *pgRep) getSelectQueryText() string {
	return `SELECT "id", 
								 "name", 
								 "created",
								 "email", 
								 "status", 								
								 "passwordHash",
								 "passwordSalt",
								 "role" 
					FROM  "public"."users" `
}

func (r *pgRep) query(sql string, values ...interface{}) ([]*user.User, error) {
	rows, err := r.pool.Query(context.Background(), sql, values...)
	if err != nil {
		return nil, err
	}

	users := make([]*user.User, 0)

	for rows.Next() {
		u := new(user.User)
		err := rows.Scan(&u.ID, &u.Name, &u.Created, &u.Email, &u.Status, &u.PasswordHash, &u.PasswordSalt, &u.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
