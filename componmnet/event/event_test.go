package event

import (
	"github.com/pleuvoir/gamine"
	"testing"
	"time"
)

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
