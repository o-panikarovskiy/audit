package pgrep

import (
	"audit/src/user"
	"database/sql"
)

func (r *pgRepository) FindByID(id string) (*user.User, error) {
	text := `SELECT id, 
								  name, 
								  created,
								  email, 
								  status, 								
								  password_hash,
								  password_salt,
								  role
					FROM  public.users
					WHERE id = $1;`

	u := &user.User{}

	err := r.db.QueryRow(text, id).Scan(
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
