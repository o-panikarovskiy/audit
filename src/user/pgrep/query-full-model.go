package pgrep

import (
	"audit/src/user"
	"context"
)

func (r *pgRepository) queryFullModel(sql string, values ...interface{}) ([]*user.User, error) {
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
