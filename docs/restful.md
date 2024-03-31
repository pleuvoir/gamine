
## restful 服务

目前对`gin`进行了简易的封装，因为是服务启动因此没有`GET`获取实例的必要。

- WithServerStartedListener 服务启动后会进行回调
- WithUseRequestLog 请求日志，使用默认的格式，需要传入`logrus`的实现
- WithGinConfig 对外暴露`gin`可以无差别的进行配置

参考以下示例，其中`WithUseRequestLog`和`WithUseRequestLog`都是可选的，示例中为了打印日志提前初始了日志组件。
在生产环境下，端口应当是传入的，建议使用获取配置文件的形式进行读取。



```go
func TestServer(t *testing.T) {
	gamine.SetWorkDir("/Users/pleuvoir/dev/space/git/gamine/test")
	gamine.SetEnvName("dev")
	gamine.InstallComponents(&log.Instance{})
	server := NewRestServer("8001")
	server.WithServerStartedListener(func(engine *Instance) {
		t.Log("启动了" + engine.port)
	})
	server.WithUseRequestLog(log.GetDefault())
	server.WithGinConfig(func(e *gin.Engine) {
		index := e.Group("/")
		{
			indexController := NewIndexController()
			index.GET("/", indexController.Index)
		}
	})
	server.Run()
}
```

输出：

```
gamine设置环境：dev
gamine从环境变量中获取到工作目录：/Users/pleuvoir/dev/space/git/gamine/test
gamine从工作目录加载应用配置文件：/Users/pleuvoir/dev/space/git/gamine/test/gamine-dev.yml
准备初始化日志: level:info, logPath:/Users/pleuvoir/dev/space/git/gamine/test/test_data/, logFilename:bak, maxAge:1440h0m0s, rotationTime:24h0m0s 
准备初始化日志: level:info, logPath:/Users/pleuvoir/dev/space/git/gamine/test/test_data/, logFilename:default, maxAge:1440h0m0s, rotationTime:24h0m0s 
restful服务已启动 @8001
    server_test.go:19: 启动了8001
```


