package helper_config

import (
	"fmt"
	"testing"
)

type Instance struct {
	Items map[string]Item `yaml:"items"`
}

type Item struct {
	Alias string `yaml:"alias"`
}

func TestParseStruct2YamlString(t *testing.T) {

	instance := Instance{Items: map[string]Item{}}

	defaultItem := Item{Alias: "default"}
	mainItem := Item{Alias: "main"}

	instance.Items["default"] = defaultItem
	instance.Items["main"] = mainItem

	t.Logf("%+v", instance)

	value, err := ParseStruct2YamlString(instance)
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}

func TestParseYamlString2Struct(t *testing.T) {
	ymlString := `items:
    default:
        alias: default
    main:
        alias: main`
	t.Log(ymlString)

	instance := &Instance{}
	err := ParseYamlString2Struct(ymlString, instance)
	if err != nil {
		panic(err)
	}
	t.Logf(fmt.Sprintf("%+v", instance))
}

func TestInjectAnotherStructByYaml(t *testing.T) {
	instance := &Instance{Items: map[string]Item{}}
	defaultItem := Item{Alias: "default"}
	mainItem := Item{Alias: "main"}
	instance.Items["default"] = defaultItem
	instance.Items["main"] = mainItem
	t.Logf("%+v", instance)

	another := &Instance2{}
	err := InjectAnotherStructByYaml(instance, another)
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", another)
}

type Instance2 struct {
	Item2 map[string]Item2 `yaml:"items"`
}

type Item2 struct {
	Alias string `yaml:"alias"`
}

func TestLoadYamlString(t *testing.T) {
	if yamlString, err := ReadYamlString("/Users/pleuvoir/dev/space/git/gamine/test/gamine-dev.yml"); err == nil {
		t.Log(yamlString)
	}
}

type AppConfig struct {
	App struct {
		Port string `yaml:"port"`
	} `yaml:"app"`
}

func TestServerWithConfig(t *testing.T) {
	path := "/Users/pleuvoir/dev/space/git/gamine/test/restful.yml"
	yamlString, _ := ReadYamlString(path)
	t.Log(yamlString)

	app := AppConfig{}
	ParseYamlString2Struct(yamlString, &app)
	t.Logf("%+v", app)

	app2 := AppConfig{}
	ParseYamlStringFromPath2Struct(path, &app2)
	t.Logf("%+v", app2)
}
