package env_test

import (
	"testing"

	"github.com/lxkedinh/stampy/env"
)

func TestLoadEnvFile(t *testing.T) {
	err := env.Load("./testdata/.env")
	if err != nil {
		t.Fatalf("Testing loading env file failed\nError: %v", err)
	}
}
