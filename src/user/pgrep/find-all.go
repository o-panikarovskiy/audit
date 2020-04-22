package pgrep

import "audit/src/user"

func (r *pgRepository) FindAll() ([]*user.User, error) {
	sql := r.getSelectModelText() + `;`
	return r.queryFullModel(sql)
}
