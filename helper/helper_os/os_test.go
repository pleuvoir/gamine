package helper_os

import (
	"path"
	"testing"
)

func TestGetEnv(t *testing.T) {
	env := GetEnvOrDefault("key", "default-key-value")
	t.Log(env)

	SetEnvQuiet("key2", "value2")

	t.Log(GetEnv("key2"))
}

func TestGetEnvOrDefault(t *testing.T) {
	env := GetEnvOrDefault("key", "default-key-value")
	t.Log(env)
}

func TestEnv(t *testing.T) {
	SetEnvQuiet("key2", "pleuvoir")
	env := GetEnv("key2")
	t.Logf(env)

	envOrDefault := GetEnvOrDefault("key", "default-key-value")
	t.Log(envOrDefault)
}

func TestGetWdQuiet(t *testing.T) {
	dir := GetWdQuiet()
	t.Log(dir)
}

func TestFileExists(t *testing.T) {
	filePath := path.Join(GetWdQuiet(), "os.go")
	t.Log(filePath)
	exists := FileExists(filePath)
	t.Log(exists)
}

func TestCurrentPath(t *testing.T) {
	currentPath, err := CurrentExecutePath()
	if err != nil {
		panic(err)
	}
	t.Log(currentPath)
}

func TestRootPath(t *testing.T) {
	t.Log(RootPath())
}

type closeImpl struct {
}

func (c *closeImpl) Close() error {
	panic("implement me")
}

func TestCloseQuietly(t *testing.T) {
	CloseQuietly(&closeImpl{})
}

func TestWaitQuit(t *testing.T) {
	WaitQuit()
}
