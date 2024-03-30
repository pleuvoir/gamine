package log

import "github.com/sirupsen/logrus"

const (
	name       = "log-engine"
	defaultLog = "default"
)

var group = make(map[string]*logrus.Logger)

type Instance struct {
	LogConfigs map[string]Config `yaml:"log-configs"`
}

type Config struct {
	Level        string `json:"level" yaml:"level"`
	Path         string `json:"path" yaml:"path"`
	Filename     string `json:"filename" yaml:"filename"`
	MaxAge       string `json:"maxAge" yaml:"maxAge"`
	RotationTime string `json:"rotationTime" yaml:"rotationTime"`
}

func (i *Instance) Run() error {
	for k, v := range i.LogConfigs {
		engine := NewEngine(&v)
		group[k] = engine.log
	}
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get(alias string) *logrus.Logger {
	log, ok := group[alias]
	if ok {
		return log
	} else {
		return nil
	}
}

func GetDefault() *logrus.Logger {
	return Get(defaultLog)
}
