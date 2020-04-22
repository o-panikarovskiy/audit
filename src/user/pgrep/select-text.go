package pgrep

func (r *pgRepository) getSelectModelText() string {
	return `SELECT "id", 
								 "name", 
								 "created",
								 "email", 
								 "status", 								
								 "passwordHash",
								 "passwordSalt",
								 "role" 
					FROM  "public"."users" `
}
