package help_config

import "gopkg.in/yaml.v3"

// ParseYamlString2Struct 解析yaml字符串为结构体
func ParseYamlString2Struct(str string, conf any) error {
	err := yaml.Unmarshal([]byte(str), conf)
	if err != nil {
		return err
	}
	return nil
}

// ParseStruct2YamlString 解析结构体为yaml字符串
func ParseStruct2YamlString(conf any) (value string, err error) {
	out, err := yaml.Marshal(conf)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// InjectAnotherStructByYaml 将原始yaml结构体的值赋值给另外一个结构体
func InjectAnotherStructByYaml(source any, another any) error {
	yamlString, err := ParseStruct2YamlString(source)
	if err != nil {
		return err
	}
	err = ParseYamlString2Struct(yamlString, another)
	if err != nil {
		return err
	}
	return nil
}
