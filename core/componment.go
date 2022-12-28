package core

import (
	"gopkg.in/yaml.v3"
)

type Component interface {
	Run() error
	GetName() string
}

var components = make(map[string]Component)

func LoadComponents(instances ...Component) {
	for _, c := range instances {
		LoadComponent(c)
	}
}

// LoadComponent 加载组件，外部可以传入组件空结构体，会从配置文件中自动映射到组件中
func LoadComponent(c Component) {
	if err := InjectComponentConfig(c.GetName(), c); err != nil {
		panic(err)
	}
	if err := c.Run(); err != nil {
		panic(err)
	}
	components[c.GetName()] = c
}

func RunComponents(instances ...Component) {
	for _, c := range instances {
		components[c.GetName()] = c
		c.Run()
	}
}

func LoadComponentFromYaml(c Component, content []byte) error {
	err := yaml.Unmarshal(content, c)
	if err != nil {
		return err
	}
	c.Run()
	return nil
}
