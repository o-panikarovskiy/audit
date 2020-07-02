package pgrep

import (
	"audit/src/user"
	"database/sql"
)

func (r *pgRepository) FindByUsername(email string) (*user.User, error) {
	text := `SELECT id, 
								  name, 
								  (extract(epoch from created)*1000) AS created,
								  email, 
								  status, 								
								  password_hash,
								  password_salt,
								  role
					FROM  public.users
					WHERE email = $1;`

	u := &user.User{}

	err := r.db.QueryRow(text, email).Scan(
		&u.ID,
		&u.Name,
		&u.Created,
		&u.Email,
		&u.Status,
		&u.PasswordHash,
		&u.PasswordSalt,
		&u.Role,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return u, nil
}
