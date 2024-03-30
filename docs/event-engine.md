


## event 事件引擎

```yaml
event-engine:
  event-configs:
    main:
      debug-log: true
      every-event-queue-size: 100
      event-type: [ "Log","Tick","Bar","Trade","Order","Asset","Position","Contract","Account","Algo","Error" ]
    timer:
      debug-log: true
      enable-timer: true
      timer-duration-second: 1
      timer-queue-size: 1000
```


`event-type`定义了事件类型，发布事件时可以指定类型，注意必须是事件类型中定义的事件。


```go
func TestRun(t *testing.T) {
    gamine.SetWorkDir("/Users/pleuvoir/dev/space/git/gamine/test")
    gamine.SetEnvName("dev")
    gamine.InstallComponents(&Instance{})

	main := Main()

	main.Register("Tick", AdaptEventHandlerFunc(func(e Event) {
		t.Log(e)
	}))

	for i := 0; i < 20; i++ {
		go func() {
			main.Put(NewEvent("Tick", "i am tick"))
		}()
	}

	timer := Timer()
	timer.Register("Timer", AdaptEventHandlerFunc(func(e Event) {
		t.Log(e)
	}))

	time.Sleep(time.Second * 50)
}
```