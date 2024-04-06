package core

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/pleuvoir/gamine/helper/helper_config"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"path/filepath"
)

type ConfigManager struct {
	fileName    string
	fileExt     string
	searchPaths []string
	config      map[string]any `yaml:",inline"`
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		config: make(map[string]any),
	}
}

func (cm *ConfigManager) GetConfig() map[string]any {
	return cm.config
}

func (cm *ConfigManager) SetConfigName(name string) {
	cm.fileName = name
}

func (cm *ConfigManager) SetConfigType(ext string) {
	cm.fileExt = ext
}

func (cm *ConfigManager) AddConfigPath(path string) {
	cm.searchPaths = append(cm.searchPaths, path)
}

func (cm *ConfigManager) LoadConfigFile() error {
	for _, dir := range cm.searchPaths {
		filePath := filepath.Join(dir, fmt.Sprintf("%s.%s", cm.fileName, cm.fileExt))
		normalizePath, _ := helper_os.NormalizePath(filePath)
		color.Warnln(fmt.Sprintf("尝试加载配置文件：%s", normalizePath))
		if err := helper_config.ParseYamlStringFromPath2Struct(normalizePath, &cm.config); err == nil {
			color.Greenln("配置文件加载成功")
			return nil
		}
	}
	return fmt.Errorf("配置文件未找到")
}
