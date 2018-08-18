package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE pokemon ADD COLUMN alola_id INTEGER")
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE pokemon DROP COLUMN alola_id")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180815160540_add_alola_id_to_pokemon", up, down, opts)
}
