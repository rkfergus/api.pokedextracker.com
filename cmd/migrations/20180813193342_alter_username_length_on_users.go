package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE users ALTER COLUMN username TYPE VARCHAR(20)")
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE users ALTER COLUMN username TYPE VARCHAR(255)")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180813193342_alter_username_length_on_users", up, down, opts)
}
