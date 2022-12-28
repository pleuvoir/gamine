package gamine

import (
	"fmt"
	"github.com/pleuvoir/gamine/core"
	"github.com/pleuvoir/gamine/helper/os"
	"path/filepath"
)

func Init() {

	//当前环境 开发、线上
	core.EnvName = os.GetEnvOrDefault(core.GamineEnv, core.Dev)

	//从环境变量中获取当前工作目录
	if core.WorkDir == "" {
		core.WorkDir = os.GetEnv(core.GamineWorkerDir)
	}

	//如果没有获取到，则使用系统工作目录
	if core.WorkDir == "" {
		core.WorkDir = os.GetWdQuiet()
		os.SetEnvQuiet(core.GamineWorkerDir, core.WorkDir)
	}

	//加载应用配置文件
	configPath := filepath.Join(core.WorkDir, fmt.Sprintf("app-%s.yml", core.EnvName))
	if err := core.LoadConfigFile(configPath); err != nil {
		panic(err)
	}

}

func InstallComponents(instances ...core.Component) {
	Init()
	core.LoadComponents(instances...)
}

func LoadComponents(instances ...core.Component) {
	core.LoadComponents(instances...)
}

func SetWorkDir(dir string) {
	core.WorkDir = dir
}
func GetWorkDir() string {
	return core.WorkDir
}

func SetEnvName(envName string) {
	core.EnvName = envName
}
func GetEnvName() string {
	return core.EnvName
}
