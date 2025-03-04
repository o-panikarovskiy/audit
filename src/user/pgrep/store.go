package pgrep

import (
	"audit/src/user"
	"time"
)

func (r *pgRepository) Store(u *user.User) (*user.User, error) {
	text := `INSERT INTO public.users(					                              
																		name, 					                             
																		email, 
																		status, 
																		password_hash,
																		password_salt,
																		role
					) 
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id, created;`

	var lastInsertID string
	var created time.Time
	err := r.db.QueryRow(text, u.Name, u.Email, u.Status, u.PasswordHash, u.PasswordSalt, u.Role).Scan(&lastInsertID, &created)
	if err != nil {
		return nil, err
	}

	u.ID = lastInsertID
	u.Created = created

	return u, nil
}
