
## 高级

一次运行多个组件，一次加载到处运行 :)

```go
func TestMore(t *testing.T) {
	logComponent := &log.Instance{LogConfigs: map[string]log.Config{}}
	logComponent.LogConfigs["test"] = log.Config{
		Level:        "debug",
		Path:         "/Users/pleuvoir/dev/space/git/gamine/test/test_data/",
		Filename:     "test",
		MaxAge:       "1440h",
		RotationTime: "24h",
	}
	gamine.InstallComponents(&hello.Instance{}, logComponent)
	t.Logf(fmt.Sprintf("%+v", hello.Get()))
	log.Get("test").Infoln("你好")
}
```

输出：

```
gamine设置环境：dev
gamine使用系统工作目录：/Users/pleuvoir/dev/space/git/gamine/test
gamine从工作目录加载应用配置文件：/Users/pleuvoir/dev/space/git/gamine/test/gamine-dev.yml
准备初始化日志: level:debug, logPath:/Users/pleuvoir/dev/space/git/gamine/test/test_data/, logFilename:test, maxAge:1440h0m0s, rotationTime:24h0m0s 
准备初始化日志: level:info, logPath:/Users/pleuvoir/dev/space/git/gamine/test/test_data/, logFilename:bak, maxAge:1440h0m0s, rotationTime:24h0m0s 
准备初始化日志: level:info, logPath:/Users/pleuvoir/dev/space/git/gamine/test/test_data/, logFilename:default, maxAge:1440h0m0s, rotationTime:24h0m0s 
    gamine_test.go:34: &{Message:hello gamine Version:1.1}
time="2024-03-30T19:37:59+08:00" level=info msg="你好"
```