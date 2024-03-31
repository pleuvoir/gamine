package core

import (
	"errors"
	"fmt"
	"github.com/pleuvoir/gamine/helper/helper_config"
)

var componentConfig = make(map[string]any)

// LoadConfigFile 加载组件配置文件
func LoadConfigFile(path string) error {
	err := helper_config.ParseYamlStringFromPath2Struct(path, &componentConfig)
	if err != nil {
		return err
	}
	return nil
}

// InjectComponentConfig 注入组件配置文件
func InjectComponentConfig(name string, conf any) error {
	if obj, ok := componentConfig[name]; ok {
		err := helper_config.InjectAnotherStructByYaml(obj, conf)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("component config not find, name:%s,conf:%+v", name, conf))
	}
}
