package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		if _, err := db.Exec("UPDATE dexes SET region = ? WHERE generation = ?", "national", 6); err != nil {
			return err
		}
		if _, err := db.Exec("UPDATE dexes SET region = ? WHERE generation = ?", "alola", 7); err != nil {
			return err
		}
		_, err := db.Exec("ALTER TABLE dexes ALTER COLUMN region SET NOT NULL")
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE dexes ALTER COLUMN region DROP NOT NULL")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180815162542_backfill_dexes_region", up, down, opts)
}
