package gamine

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/pleuvoir/gamine/core"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"os"
)

// InstallComponents 安装，运行组件，该方法只应该被调用一次
func InstallComponents(instances ...core.IComponent) {
	config := loadConfig()
	core.LoadComponents(config, instances...)
}

// RunComponents 运行组件，适用于无需加载配置文件的场景，无需调用 InstallComponents
func RunComponents(instances ...core.IComponent) {
	core.RunComponents(instances...)
}

func loadConfig() map[string]any {
	//当前环境 开发、线上
	env := GetEnvName()
	if env == "" {
		SetEnvName(core.Dev)
	}
	color.Greenln(fmt.Sprintf("gamine设置环境：%s", core.EnvName))

	//切换工作目录
	workDir := GetWorkDir()
	if workDir == "" {
		panic("工作目录为空，请设置")
	}
	normalizePath, err := helper_os.NormalizePath(workDir)
	if err != nil {
		panic(fmt.Sprintf("获取工作目录绝对路径失败，%s", err))

	}
	workDir = normalizePath
	if err := os.Chdir(workDir); err != nil {
		panic(fmt.Sprintf("切换工作目录失败，%s", err))
	}

	color.Greenln(fmt.Sprintf("gamine已切换到到工作目录：%s", workDir))

	//从多目录加载应用配置文件
	configManager := core.NewConfigManager()
	configManager.SetConfigName(fmt.Sprintf("gamine-%s", GetEnvName()))
	configManager.SetConfigType("yml")
	configManager.AddConfigPath(".")
	configManager.AddConfigPath("./bin")
	configManager.AddConfigPath("../bin")
	configManager.AddConfigPath("./configs")
	configManager.AddConfigPath("../configs")

	if err := configManager.LoadConfigFile(); err != nil {
		color.Redln(fmt.Sprintf("gamine从工作目录加载配置文件失败: %s", err))
		panic(err)
	}
	return configManager.GetConfig()
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
