package helper_os

import (
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

// GetEnvOrDefault 获取环境变量，未获取到返回给定的默认值
func GetEnvOrDefault(key string, defaultVal string) string {
	v := GetEnv(key)
	if v == "" {
		return defaultVal
	}
	return v
}

// GetEnv 获取环境变量
func GetEnv(key string) string {
	return os.Getenv(key)
}

// SetEnvQuiet 安静的设置环境变量
func SetEnvQuiet(key, value string) {
	_ = os.Setenv(key, value)
}

// GetWdQuiet 安静的获取工作目录 （示例）
func GetWdQuiet() (dir string) {
	dir, _ = os.Getwd()
	return dir
}

// FileExists 文件是否存在
func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	if info == nil {
		return false
	}
	return true
}

// CurrentExecutePath 获取当前的执行文件所在的目录
func CurrentExecutePath() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(dir), nil
}

// RootPath 获取项目根路径 （示例）
func RootPath() (string, error) {
	dir, err := filepath.Abs("")
	if err != nil {
		return "", err
	}
	return filepath.Dir(dir), nil
}

// Abs 获取绝对路径 （示例）
func Abs(path string) string {
	if abs, err := filepath.Abs(path); err == nil {
		return abs
	}
	return path
}

// CloseQuietly 安静的调用Close()
func CloseQuietly(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

// WaitQuit 同步阻塞等待退出
func WaitQuit() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
