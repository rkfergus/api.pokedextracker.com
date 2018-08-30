package main

import (
	"log"
	"os"

	"github.com/pokedextracker/api.pokedextracker.com/pkg/config"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/database"
	"github.com/robinjoseph08/go-pg-migrations"
)

const directory = "./cmd/migrations"

func main() {
	// log := logger.New()

	cfg := config.New()
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalln(err)
		// log.Err(err).Fatal("database error")
	}
	// db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
	// 	query, err := event.FormattedQuery()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	log.Println(query)
	// 	// log.Info("processed query", logger.Data{"time": time.Since(event.StartTime), "query": query})
	// })

	err = migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
		// log.Err(err).Fatal("migration error")
	}
}
