package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		if _, err := db.Exec("ALTER TABLE evolutions DROP CONSTRAINT evolutions_stone_check"); err != nil {
			return err
		}
		if _, err := db.Exec("ALTER TABLE evolutions ADD CONSTRAINT evolutions_stone_check CHECK (stone = ANY (ARRAY['fire'::text, 'water'::text, 'thunder'::text, 'leaf'::text, 'moon'::text, 'sun'::text, 'shiny'::text, 'dusk'::text, 'dawn'::text, 'ice'::text])) NOT VALID"); err != nil {
			return err
		}
		_, err := db.Exec("ALTER TABLE evolutions VALIDATE CONSTRAINT evolutions_stone_check")
		return err
	}

	down := func(db orm.DB) error {
		if _, err := db.Exec("ALTER TABLE evolutions DROP CONSTRAINT evolutions_stone_check"); err != nil {
			return err
		}
		if _, err := db.Exec("ALTER TABLE evolutions ADD CONSTRAINT evolutions_stone_check CHECK (stone = ANY (ARRAY['fire'::text, 'water'::text, 'thunder'::text, 'leaf'::text, 'moon'::text, 'sun'::text, 'shiny'::text, 'dusk'::text, 'dawn'::text])) NOT VALID"); err != nil {
			return err
		}
		_, err := db.Exec("ALTER TABLE evolutions VALIDATE CONSTRAINT evolutions_stone_check")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180815180538_add_ice_stone_to_evolutions_stone_enum", up, down, opts)
}
