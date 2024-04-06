
## socketIO

绑定了`request`和`push`两种事件，使用`gin`做为引擎异步启动。需要启动后手动阻塞进程。

- WithRequest 处理`request`事件

```go

// 用这个测试 https://amritb.github.io/socketio-client-tool/v1/  ws://localhost:8000
func TestNewSocketServer(t *testing.T) {

	goSocketIO := New(8000)
	_ = goSocketIO.WithRequest(func(msg RequestMessage) ResponseMessage {
		t.Log(fmt.Sprintf("%+v", msg))
		return ResponseMessage{Data: "hello world"}
	})

	_ = goSocketIO.Run()

    go func() {
        count := 1
        for true {
			//可以随时获取
            Get().PushResponse(ResponseMessage{Data: count, MethodName: "updateCount"})
            count++
            time.Sleep(3 * time.Second)
        }
    }()

	//阻塞当前进程
	helper_os.WaitQuit()

	t.Log("ok")
}
```


输出：

```
socketIO已启动 @8000
```




