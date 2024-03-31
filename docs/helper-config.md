

## 配置文件相关工具类

### ParseStruct2YamlString 将结构体转换为yml字符串   

```go
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
```

输出：

```
items:
    default:
        alias: default
    main:
        alias: main
```

### ParseYamlString2Struct 将yml字符串映射到结构体 

```go
func TestParseYamlString2Struct(t *testing.T) {
	ymlString := `items:
    default:
        alias: default
    main:
        alias: main`

	instance := &Instance{}
	err := ParseYamlString2Struct(ymlString, instance)
	if err != nil {
		panic(err)
	}
	t.Logf(fmt.Sprintf("%+v", instance))
}
```

输出：

```
&{Items:map[default:{Alias:default} main:{Alias:main}]}
```

### InjectAnotherStructByYaml 将原始yaml结构体的值赋值给另外一个结构体

这个方法类似于copy对象，工作原理是先将一个结构体转为yaml字符串，然后将该字符串再反序列化到另一个结构体上。
因此两个结构体可以不相同，但是必要字段只要用 tag `yaml:"?"` 声明也可以完成映射。



```go
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
```

输出：

```
&{Items:map[default:{Alias:default} main:{Alias:main}]}
&{Item2:map[default:{Alias:default} main:{Alias:main}]}
```

### ReadYamlString 读取yml文件转换为字符串

```go
func TestLoadYamlString(t *testing.T) {
	if yamlString, err := ReadYamlString("/Users/pleuvoir/dev/space/git/gamine/test/gamine-dev.yml"); err == nil {
		t.Log(yamlString)
	}
}
```

输出：

```yaml
hello:
  message: "hello gamine"
  version: 1.1
```


### ParseYamlStringFromPath2Struct 从文件路径读取yaml字符串到结构体

```yaml
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
```

输出：

```
config_test.go:92: app:
      port: "8081"
    
    
config_test.go:96: {App:{Port:8081}}
config_test.go:101: {App:{Port:8081}}
```