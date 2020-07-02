package pgrep

import "audit/src/user"

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
					RETURNING id;`

	lastInsertID := ""
	err := r.db.QueryRow(text, u.Name, u.Email, u.Status, u.PasswordHash, u.PasswordSalt, u.Role).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}

	u.ID = lastInsertID
	return u, nil
}
