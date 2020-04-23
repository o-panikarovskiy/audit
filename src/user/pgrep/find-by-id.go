package pgrep

import "audit/src/user"

func (r *pgRepository) FindByID(id string) (*user.User, error) {
	sql := r.getSelectModelText() + ` WHERE "id" = $1 LIMIT 1;`

	res, err := r.queryFullModel(sql, id)
	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		return res[0], nil
	}

	return nil, nil
}
