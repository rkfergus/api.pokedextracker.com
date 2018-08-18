package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	const limit = 1000
	type user struct {
		tableName struct{} `sql:"users,alias:users"`

		ID int
	}
	type dex struct {
		tableName struct{} `sql:"dexes,alias:dexes"`

		ID         string
		UserID     int
		Title      string
		Slug       string
		Shiny      bool `sql:",notnull"`
		Generation int
	}

	batch := func(db orm.DB) (int, error) {
		users := []user{}
		dexes := []dex{}

		if err := db.Model(&users).Join("LEFT OUTER JOIN dexes ON users.id = dexes.user_id").Where("dexes.id IS NULL").Limit(limit).Select(); err != nil {
			return 0, err
		}

		// go-pg returns an error when inserting an empty slice
		if len(users) == 0 {
			return 0, nil
		}

		for _, u := range users {
			dexes = append(dexes, dex{
				UserID:     u.ID,
				Title:      "Living Dex",
				Slug:       "living-dex",
				Shiny:      false,
				Generation: 6,
			})
		}

		if _, err := db.Model(&dexes).Insert(); err != nil {
			return 0, err
		}

		return len(dexes), nil
	}

	up := func(db orm.DB) error {
		for {
			count, err := batch(db)
			if err != nil {
				return err
			}
			if count != limit {
				break
			}
		}
		return nil
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DELETE FROM dexes WHERE title = ?", "Living Dex")
		return err
	}

	opts := migrations.MigrationOptions{DisableTransaction: true}

	migrations.Register("20180813201845_backfill_dexes_table", up, down, opts)
}
