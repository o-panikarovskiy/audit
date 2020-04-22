package pgrep

import "audit/src/user"

func (r *pgRepository) Store(u *user.User) (*user.User, error) {
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

	res, err := r.queryFullModel(sql, u.Name, u.Email, u.Status, u.PasswordHash, u.PasswordSalt, u.Role)

	if err != nil {
		return nil, err
	}

	return res[0], nil
}
