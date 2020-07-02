package pgrep

import (
	"audit/src/user"
)

func (r *pgRepository) FindAll() ([]*user.User, error) {
	text := `SELECT id, 
								  name, 
								  (extract(epoch from created)*1000) AS created,
								  email, 
								  status, 								
								  password_hash,
								  password_salt,
								  role
					FROM  public.users;`

	rows, err := r.db.Query(text)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*user.User, 0)
	for rows.Next() {
		u := &user.User{}

		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Created,
			&u.Email,
			&u.Status,
			&u.PasswordHash,
			&u.PasswordSalt,
			&u.Role,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
