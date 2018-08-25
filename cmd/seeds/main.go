package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/go-pg/pg"
	"github.com/pokedextracker/api.pokedextracker.com/config"
	"github.com/pokedextracker/api.pokedextracker.com/database"
)

const directory = "./cmd/seeds"

func main() {
	cfg := config.New()
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	if cfg.Environment == "staging" || cfg.Environment == "production" {
		fmt.Printf("Cannot seed database in %s\n", cfg.Environment)
	}

	loadFile(db, "truncate.sql")
	loadFile(db, "game_families.sql")
	loadFile(db, "games.sql")
}

func loadFile(db *pg.DB, file string) {
	path := filepath.Join(directory, file)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Running %q...\n", path)
	_, err = db.Exec(string(b))
	if err != nil {
		log.Fatalln(err)
	}
}
