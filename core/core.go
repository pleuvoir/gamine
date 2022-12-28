package core

var (
	WorkDir       string
	EnvName       string
	deferHandlers []DeferHandle
)

const (
	Dev             = "dev"
	Prod            = "prod"
	GamineEnv       = "gamine_env"
	GamineWorkerDir = "gamine_worker_dir"
)

type DeferHandle func()

func DeferRun() {
	for _, f := range deferHandlers {
		f()
	}
}

func AddDefer(handler DeferHandle) {
	deferHandlers = append(deferHandlers, handler)
}
