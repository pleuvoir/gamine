package os

import (
	"os"
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

// GetWdQuiet 安静的获取工作目录
func GetWdQuiet() (dir string) {
	dir, _ = os.Getwd()
	return dir
}

// FileExists 文件是否存在
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if info == nil {
		return false
	}
	return true
}
