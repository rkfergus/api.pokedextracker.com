package database

import (
	"testing"

	"github.com/pokedextracker/api.pokedextracker.com/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := config.New()

	db, err := New(cfg)

	assert.Nil(t, err)
	assert.NotNil(t, db)

	cfg.DatabaseName = "bad_db"

	_, err = New(cfg)

	assert.Error(t, err, "expected error when connection fails")
}
