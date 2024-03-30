package helper_os

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	env := GetEnvOrDefault("key", "default-key-value")
	t.Log(env)

	SetEnvQuiet("key2", "value2")

	t.Log(GetEnv("key2"))
}
