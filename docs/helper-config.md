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