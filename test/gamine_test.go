package test

import (
	"fmt"
	"github.com/pleuvoir/gamine"
	"github.com/pleuvoir/gamine/component/hello"
	"github.com/pleuvoir/gamine/component/log"
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
