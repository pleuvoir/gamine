package event

const name = "event-engine"

type Instance struct {
	EventConfigs map[string]Config `yaml:"event-configs"`
}

type Config struct {
	EventType           []string `yaml:"event-type"`
	EnableTimer         bool     `yaml:"enable-timer"`
	TimerQueueSize      int      `yaml:"timer-queue-size"`
	EveryEventQueueSize int      `yaml:"every-event-queue-size"`
	DebugLog            bool     `yaml:"debug-log"`
}

var group = make(map[string]*Engine)

func (i *Instance) Run() error {
	for k, v := range i.EventConfigs {
		//启动事件引擎
		engine := NewEventEngine(k, v.EventType, v.EveryEventQueueSize, v.EnableTimer, v.TimerQueueSize)
		engine.DebugLog = v.DebugLog
		if v.EnableTimer {
			engine.StartAll()
		} else {
			engine.startConsumer()
		}
		group[k] = engine
	}
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get(alias string) *Engine {
	engine, ok := group[alias]
	if ok {
		return engine
	} else {
		return nil
	}
}

func Main() *Engine {
	return Get("main")
}

func Timer() *Engine {
	return Get("timer")
}
