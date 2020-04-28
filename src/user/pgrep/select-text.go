package pgrep

func (r *pgRepository) getSelectModelText() string {
	return `SELECT "id", 
								 "name", 
								 (extract(epoch from "created")*1000) AS "created",
								 "email", 
								 "status", 								
								 "passwordHash",
								 "passwordSalt",
								 "role" 
					FROM  "public"."users" `
}
