package log

import (
	"github.com/pleuvoir/gamine"
	"testing"
)

func TestLoad(t *testing.T) {
	gamine.SetWorkDir("/Users/pleuvoir/dev/space/git/gamine/test")
	gamine.SetEnvName("dev")
	gamine.InstallComponents(&Instance{})
	GetDefault().Infoln("default work")
	Get("bak").Infoln("bak work")
}

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
