package helper_config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"gopkg.in/yaml.v3"
	"os"
)

// ReadYamlString 读取yaml字符串
func ReadYamlString(filePath string) (string, error) {
	if !helper_os.FileExists(filePath) {
		return "", errors.New(fmt.Sprintf("文件不存在，%s", filePath))
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ParseYamlStringFromPath2Struct 从文件路径读取yaml字符串到结构体
func ParseYamlStringFromPath2Struct(filePath string, conf any) error {
	yamlString, err := ReadYamlString(filePath)
	if err != nil {
		return err
	}
	err = ParseYamlString2Struct(yamlString, conf)
	if err != nil {
		return err
	}
	return nil
}

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
