package application

import (
	"testing"

	"github.com/go-pg/pg"
	"github.com/pokedextracker/api.pokedextracker.com/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	db := &pg.DB{}
	cfg := &config.Config{}

	app := New(db, cfg)
	assert.NotNil(t, app, "returned app shouldn't be nil")
}
