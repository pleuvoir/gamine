package hello

import (
	"fmt"
	"github.com/pleuvoir/gamine"
	"testing"
)

func TestRun(t *testing.T) {
	gamine.SetEnvName("dev")
	gamine.SetWorkDir("/Users/pleuvoir/dev/space/git/gamine/test")
	gamine.InstallComponents(&Instance{})
	instance := Get()
	t.Logf(fmt.Sprintf("%+v", instance))
}
