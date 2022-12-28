package os

import (
	"github.com/pleuvoir/gamine/helper/os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	env := os.GetEnvOrDefault("key", "default-key-value")
	t.Log(env)

	os.SetEnvQuiet("key2", "value2")

	t.Log(os.GetEnv("key2"))
}
