package core

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var componentConfig = make(map[string]any)

// LoadConfigFile 加载组件配置文件
func LoadConfigFile(path string) error {
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
		marshal, err := yaml.Marshal(obj)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(marshal, conf)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("component config not find, name:%s,conf:%+v", name, conf))
	}
}
