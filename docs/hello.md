


## hello 组件


hello 是一个示例组件，旨在阐述组件的运行机制。

组件的实现如下：

```go
package hello

type Instance struct {
	Message string  `yaml:"message"`
	Version float64 `yaml:"version"`
}

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	return nil
}

func (i *Instance) GetName() string {
	return "hello"
}

func Get() *Instance {
	return instance
}
```

对应的它其实实现了`IComponent`接口。具体怎么获取该组件取决于自己的实现，一般而言会命名为`Get`。组件结构体命名为`Instance`。

```go
type IComponent interface {
	Run() error
	GetName() string
}
```


```go
func TestRun(t *testing.T) {
    gamine.SetEnvName("dev")
    gamine.SetWorkDir("/Users/pleuvoir/dev/space/git/gamine/test")
    gamine.InstallComponents(&Instance{})
    instance := Get()
    t.Logf(fmt.Sprintf("%+v", instance))
}
```

接着你需要在`dev.yml`文件中配置如下的内容。组件加载后即可自动映射到对应的结构属性中。

```yaml
hello:
  message: "hello gamine"
  version: 1.1
```


输出：

```
gamine设置环境：dev
gamine从环境变量中获取到工作目录：/Users/pleuvoir/dev/space/git/gamine/test
gamine从工作目录加载应用配置文件：/Users/pleuvoir/dev/space/git/gamine/test/gamine-dev.yml
hello_test.go:13: &{Message:hello gamine Version:1.1}
```


当然了，抛开自动映射`yml`的机制，`gamine`也提供了灵活的编码形式。参考如下的代码：

```go
func TestRun(t *testing.T) {
	i := &Instance{LogConfigs: map[string]Config{}}
	i.LogConfigs["test"] = Config{
		Level:        "debug",
		Path:         "/Users/pleuvoir/dev/space/git/gamine/test/test_data/",
		Filename:     "test",
		MaxAge:       "1440h",
		RotationTime: "24h",
	}
	gamine.RunComponents(i)

	Get("test").Infoln("test load")
}
```

我们通过手动构建完整的组件结构体，`RunComponents`完成组件的运行。