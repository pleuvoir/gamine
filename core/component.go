package core

import (
	"github.com/pleuvoir/gamine/helper/helper_config"
	"gopkg.in/yaml.v3"
)

type IComponent interface {
	Run() error
	GetName() string
}

var components = make(map[string]IComponent)

func LoadComponents(componentConf map[string]any, instances ...IComponent) {
	for _, c := range instances {
		if conf, ok := componentConf[c.GetName()]; ok {
			LoadComponent(c, conf)
		}
	}
}

// LoadComponent 加载组件，外部可以传入组件空结构体，会从配置文件中自动映射到组件中
func LoadComponent(c IComponent, conf any) {
	if err := helper_config.InjectAnotherStructByYaml(conf, c); err != nil {
		panic(err)
	}
	if err := c.Run(); err != nil {
		panic(err)
	}
	components[c.GetName()] = c
}

func RunComponents(instances ...IComponent) {
	for _, c := range instances {
		components[c.GetName()] = c
		c.Run()
	}
}

func RunComponentFromYaml(c IComponent, content []byte) error {
	err := yaml.Unmarshal(content, c)
	if err != nil {
		return err
	}
	c.Run()
	return nil
}
