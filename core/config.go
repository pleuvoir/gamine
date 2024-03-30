package core

import (
	"errors"
	"fmt"
	"github.com/pleuvoir/gamine/helper/help_config"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"gopkg.in/yaml.v3"
	"os"
)

var componentConfig = make(map[string]any)

// LoadConfigFile 加载组件配置文件
func LoadConfigFile(path string) error {
	if !helper_os.FileExists(path) {
		return errors.New(fmt.Sprintf("文件不存在，%s", path))
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &componentConfig)
	if err != nil {
		return err
	}
	return nil
}

// InjectComponentConfig 注入组件配置文件
func InjectComponentConfig(name string, conf any) error {
	if obj, ok := componentConfig[name]; ok {
		err := help_config.InjectAnotherStructByYaml(obj, conf)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("component config not find, name:%s,conf:%+v", name, conf))
	}
}
