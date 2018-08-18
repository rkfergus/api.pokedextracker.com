package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE users ADD COLUMN stripe_id TEXT")
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE users DROP COLUMN stripe_id")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180817133802_add_stripe_id_to_users", up, down, opts)
}
