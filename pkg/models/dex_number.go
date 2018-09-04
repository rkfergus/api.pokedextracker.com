package models

// GameFamilyDexNumber is a model representing the connection between a Pokemon
// and a GameFamily.
type GameFamilyDexNumber struct {
	GameFamilyID string `pg:",pk"`
	PokemonID    int    `pg:",pk"`
	DexNumber    int
}
