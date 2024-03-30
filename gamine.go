package gamine

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/pleuvoir/gamine/core"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"path/filepath"
)

func Init() {

	//当前环境 开发、线上
	core.EnvName = helper_os.GetEnvOrDefault(core.GamineEnv, core.Dev)
	color.Greenln(fmt.Sprintf("gamine设置环境：%s", core.EnvName))

	//从环境变量中获取当前工作目录
	if core.WorkDir == "" {
		core.WorkDir = helper_os.GetEnv(core.GamineWorkerDir)
	}

	//如果没有获取到，则使用系统工作目录
	if core.WorkDir == "" {
		core.WorkDir = helper_os.GetWdQuiet()
		helper_os.SetEnvQuiet(core.GamineWorkerDir, core.WorkDir)
		color.Greenln(fmt.Sprintf("gamine使用系统工作目录：%s", core.WorkDir))
	} else {
		color.Greenln(fmt.Sprintf("gamine从环境变量中获取到工作目录：%s", core.WorkDir))
	}

	//从工作目录加载应用配置文件
	configPath := filepath.Join(core.WorkDir, fmt.Sprintf("gamine-%s.yml", core.EnvName))
	color.Greenln(fmt.Sprintf("gamine从工作目录加载应用配置文件：%s", configPath))

	if err := core.LoadConfigFile(configPath); err != nil {
		color.Redln(fmt.Sprintf("gamine从工作目录加载应用配置文件失败: %s", err))
		panic(err)
	}

}

// InstallComponents 安装，运行组件，该方法只应该被调用一次
func InstallComponents(instances ...core.IComponent) {
	Init()
	core.LoadComponents(instances...)
}

// RunComponents 运行组件，适用于无需加载配置文件的场景，无需调用 InstallComponents
func RunComponents(instances ...core.IComponent) {
	core.RunComponents(instances...)
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
