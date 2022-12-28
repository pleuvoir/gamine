package test

import (
	"fmt"
	"github.com/pleuvoir/gamine"
	"github.com/pleuvoir/gamine/componmnet/hello"
	"testing"
)

func TestRun(t *testing.T) {
	gamine.InstallComponents(&hello.Instance{})
	instance := hello.Get()
	t.Logf(fmt.Sprintf("%+v", instance))
}

func TestWorkDir(t *testing.T) {
	gamine.SetWorkDir("./bin")
	t.Logf(gamine.GetWorkDir())
}
