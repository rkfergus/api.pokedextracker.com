package test

import (
	"testing"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/require"
)

// TruncateTables truncates many of the tables in the database to be able to
// start from scratch for a test.
func TruncateTables(t *testing.T, db *pg.DB) {
	t.Helper()

	_, err := db.Exec(`
		TRUNCATE dexes CASCADE;
		TRUNCATE users CASCADE;
	`)
	require.NoError(t, err)
}
