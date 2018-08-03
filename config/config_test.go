package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	cfg := New()
	assert.NotNil(t, cfg, "returned config shouldn't be nil")
}

func TestEnvironments(t *testing.T) {
	originalEnv := os.Getenv(environmentENV)
	defer func() {
		err := os.Setenv(environmentENV, originalEnv)
		require.Nil(t, err, "unexpected error restoring original environment")
	}()

	envs := []string{"development", "staging", "production"}

	for _, env := range envs {
		err := os.Setenv(environmentENV, env)
		require.Nil(t, err, "unexpected error overwriting environment")

		cfg := New()
		assert.Equal(t, cfg.Environment, env, "incorrect environment")
	}
}
