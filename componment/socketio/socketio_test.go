package socketio

import (
	"fmt"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"testing"
	"time"
)

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
