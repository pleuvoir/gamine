package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/pleuvoir/gamine"
	"github.com/pleuvoir/gamine/componment/log"
	"github.com/pleuvoir/gamine/helper/helper_config"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

type AppConfig struct {
	App struct {
		Port string `yaml:"Port"`
	} `yaml:"app"`
}

func TestServerWithConf(t *testing.T) {

	path := "/Users/pleuvoir/dev/space/git/gamine/test/restful.yml"

	app := &AppConfig{}
	helper_config.ParseYamlStringFromPath2Struct(path, app)

	server := NewRestServer(app.App.Port)
	server.WithServerStartedListener(func(engine *Instance) {
		t.Log("启动了" + engine.Port)
	})
	server.WithCors()
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

func TestServer(t *testing.T) {
	gamine.SetWorkDir("../../test/")
	gamine.SetEnvName("dev")
	gamine.InstallComponents(&log.Instance{})
	server := NewRestServer("8001")
	server.WithServerStartedListener(func(engine *Instance) {
		t.Log("启动了" + engine.Port)
	})
	server.WithCors()
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

func Router(engine *gin.Engine) {
	//executePath, err := currentExecutePath()
	//if err != nil {
	//	panic(err)
	//}
	//frontendPath := filepath.Join(executePath, "../frontend")
	//if !helper_os.FileExists(frontendPath) {
	//	color.Yellowf("gin静态文件打包路径，可执行路径下未找到 %s", frontendPath)
	//	color.Yellowln()
	//	rootPath, _ := helper_os.RootPath()
	//	frontendPath = filepath.Join(rootPath, "frontend")
	//	color.Yellowf("尝试使用项目根路径..%s", frontendPath)
	//	color.Yellowln()
	//}
	//
	//if !helper_os.FileExists(frontendPath) {
	//	color.Redln("gin启动失败，未找到静态资源文件路径。")
	//	panic("gin启动失败，未找到静态资源文件路径")
	//}
	//
	//color.Greenf("gin静态文件打包路径 %s", frontendPath)
	//color.Greenln()
	//engine.StaticFS("app", http.Dir(frontendPath))
	//engine.NoRoute(func(c *gin.Context) {
	//	c.File(filepath.Join(frontendPath, "index.html"))
	//})
	index := engine.Group("/")
	{
		indexController := NewIndexController()
		index.GET("/", indexController.Index)
	}
}

func currentExecutePath() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(dir), nil
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

type IndexController struct {
}

func (t *IndexController) Index(ct *gin.Context) {
	ct.JSON(http.StatusOK, map[string]string{"name": "pleuvoir"})
}
