package pgrep

import (
	"audit/src/user"
)

func (r *pgRepository) FindByUsername(email string) (*user.User, error) {
	sql := r.getSelectModelText() + ` WHERE "email" = $1 LIMIT 1;`

	res, err := r.queryFullModel(sql, email)
	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		return res[0], nil
	}

	return nil, nil
}
