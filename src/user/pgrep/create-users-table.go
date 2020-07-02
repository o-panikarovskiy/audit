package pgrep

import (
	"log"
)

func (r *pgRepository) createUsersTable() error {
	log.Println("Create users table if not exists...")

	sql := `
    CREATE TABLE IF NOT EXISTS public.users (
      id               uuid DEFAULT gen_random_uuid() PRIMARY KEY,
      name         		 varchar(100),
      created          timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc'),
      email            varchar(256) NOT NULL UNIQUE,
      status           varchar(20) NOT NULL DEFAULT 'unverified',      
      password_hash    varchar(256),
      password_salt    varchar(256),
      role             smallint NOT NULL DEFAULT 2
    );

    CREATE UNIQUE INDEX IF NOT EXISTS users_lower_case_email ON public.users ((lower(email)));
  `
	_, err := r.db.Exec(sql)

	return err
}
