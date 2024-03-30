package hello

type Instance struct {
	Message string  `yaml:"message"`
	Version float64 `yaml:"version"`
}

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	return nil
}

func (i *Instance) GetName() string {
	return "hello"
}

func Get() *Instance {
	return instance
}
