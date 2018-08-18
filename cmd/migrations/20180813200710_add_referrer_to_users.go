package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE users ADD COLUMN referrer TEXT")
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE users DROP COLUMN referrer")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180813200710_add_referrer_to_users", up, down, opts)
}
