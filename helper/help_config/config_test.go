package help_config

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
