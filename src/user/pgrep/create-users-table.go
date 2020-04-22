package pgrep

import "context"

func (r *pgRep) createUsersTable() error {
	sql := `
    CREATE TABLE IF NOT EXISTS "public"."users" (
      "id"              uuid DEFAULT gen_random_uuid() PRIMARY KEY,
      "name"         		varchar(100),
      "created"         timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc'),
      "email"           varchar(256) NOT NULL UNIQUE,
      "status"          varchar(20) NOT NULL DEFAULT 'unverified',      
      "passwordHash"    varchar(256),
      "passwordSalt"    varchar(256),
      "role"            smallint NOT NULL DEFAULT 2
    );

    CREATE UNIQUE INDEX IF NOT EXISTS users_lower_case_email ON "public"."users" ((lower("email")));
  `

	_, err := r.pool.Exec(context.Background(), sql)

	return err
}
