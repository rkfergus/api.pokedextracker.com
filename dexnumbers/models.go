package dexnumbers

type GameFamilyDexNumber struct {
	GameFamilyID string `gorm:"primary_key"`
	PokemonID    int    `gorm:"primary_key"`
	DexNumber    int
}
