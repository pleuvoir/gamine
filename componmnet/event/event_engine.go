package event

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/panjf2000/ants"
	"sync"
	"time"
)

// Engine 事件引擎
type Engine struct {
	Name           string
	ConsumerActive bool
	TimerActive    bool
	TimeDuration   time.Duration
	HandlersMap    map[string][]eventListener
	TimerEventChan chan Event
	stopChan       chan struct{}
	registerMutex  sync.Mutex
	queues         map[string]*eventQueue
	DebugLog       bool
}

// 事件处理器
type eventListener interface {
	Process(event Event)
}

// AdaptEventHandlerFunc 接口适配器
// 可以将函数原型为 fun(event Event)的函数直接做为eventHandler接口的实现进行传入
// 可以无需定义结构体
type AdaptEventHandlerFunc func(e Event)

func (funcW AdaptEventHandlerFunc) Process(event Event) {
	funcW(event)
}

type Event struct {
	EventType string
	EventData any
}

// NewEvent 新事件
func NewEvent(eventType string, data any) Event {
	return Event{eventType, data}
}

// 每个事件使用单独的队列
type eventQueue struct {
	ch   chan Event
	t    string
	pool *ants.Pool
}

func (q *eventQueue) shutdown() {
	defer func() {
		close(q.ch)
		q.pool.Release()
	}()
}

func (q *eventQueue) send(e Event) {
	q.ch <- e
}

func newEventQueue(t string, everyEventQueueSize int) *eventQueue {
	pool, _ := ants.NewPool(everyEventQueueSize) //可以控制有多少协程并发处理任务
	return &eventQueue{ch: make(chan Event), t: t, pool: pool}
}

// NewEventEngine 创建引擎
func NewEventEngine(name string, eventType []string, everyEventQueueSize int, enableTimer bool, timerQueueSize int) *Engine {
	engine := Engine{
		ConsumerActive: false,
		HandlersMap:    map[string][]eventListener{},
		stopChan:       make(chan struct{}),
		queues:         map[string]*eventQueue{},
		Name:           name,
	}

	if enableTimer {
		engine.TimeDuration = time.Second
		engine.TimerEventChan = make(chan Event, timerQueueSize)
	}

	for _, t := range eventType {
		engine.queues[t] = newEventQueue(t, everyEventQueueSize)
	}
	return &engine
}

// Process 处理事件
func (e *Engine) Process(event Event) {
	eventHandlers := e.HandlersMap[event.EventType]
	for _, handler := range eventHandlers {
		handler.Process(event)
	}
}

// Register 注册事件处理器
func (e *Engine) Register(eventType string, handler eventListener) {
	e.registerMutex.Lock()
	defer e.registerMutex.Unlock()
	HandlersMap := e.HandlersMap
	eventHandlers := HandlersMap[eventType]
	eventHandlers = append(eventHandlers, handler)
	HandlersMap[eventType] = eventHandlers
}

// UnRegister 取消事件处理器
func (e *Engine) UnRegister(eventType string, handler eventListener) {
	e.registerMutex.Lock()
	defer e.registerMutex.Unlock()
	handlersMap := e.HandlersMap
	eventHandlers := handlersMap[eventType]
	var newHandlers []eventListener
	for _, cur := range eventHandlers {
		if cur == handler {
			continue
		}
		newHandlers = append(newHandlers, cur)
	}
	handlersMap[eventType] = newHandlers
	//没有处理器则将这个类型移除
	if len(eventHandlers) == 0 {
		delete(handlersMap, eventType)
	}
}

// StopAll 停止所有
func (e *Engine) StopAll() {
	e.stopConsumer()
	e.stopTimer()
}

// 停止周期引擎
func (e *Engine) stopTimer() {
	if e.TimerActive {
		e.TimerActive = false
		e.stopChan <- struct{}{}
	}
}

// 停止普通事件引擎
func (e *Engine) stopConsumer() {
	if e.ConsumerActive {
		e.ConsumerActive = false
		for _, e := range e.queues {
			e.shutdown()
		}
	}
}

// StartAll 启动所有
func (e *Engine) StartAll() {
	e.startTimer()
	e.startConsumer()
}

// 消费消息
func (e *Engine) startConsumer() {
	if e.ConsumerActive {
		return
	}
	for _, eq := range e.queues {
		go func(q *eventQueue) {
		over:
			for e.ConsumerActive {
				select {
				case event, ok := <-q.ch:
					if !ok {
						if e.DebugLog {
							color.Redln(fmt.Sprintf("[%s][%s]事件引擎接收到关闭信号，终止事件监听。", e.Name, q.t))
						}
						break over
					}
					err := q.pool.Submit(func() {
						if e.DebugLog {
							color.Greenln(fmt.Sprintf("[%s][%s]事件引擎提交事件到协程池中执行。", e.Name, q.t))
						}
						e.Process(event)
					})
					if err != nil {
						if e.DebugLog {
							color.Redln(fmt.Sprintf("[%s][%s]事件引擎在协程池中处理任务失败。err=%s", e.Name, q.t, err))
						}
					}
				}
			}
		}(eq)
	}
	e.ConsumerActive = true
	if e.DebugLog {
		color.Greenln(fmt.Sprintf("[%s]事件引擎普通消费者已启动。", e.Name))
	}
}

func (e *Engine) startTimer() {
	if !e.TimerActive {
		e.startTimerProducer()
		e.startTimerConsumer()
	}
	e.TimerActive = true
}

// 启动定时器消费者
func (e *Engine) startTimerConsumer() {
	go func() {
	outer:
		for {
			select {
			case event, ok := <-e.TimerEventChan:
				if !ok {
					if e.DebugLog {
						color.Redln(fmt.Sprintf("[%s]事件引擎定时器消费者接收到关闭信号，已终止事件监听。", e.Name))
					}
					break outer
				}
				e.Process(event)
			}
		}
	}()
	if e.DebugLog {
		color.Greenln(fmt.Sprintf("[%s]事件引擎定时发布消费者已启动。", e.Name))
	}
}

// 启动定时器生产者，周期执行
func (e *Engine) startTimerProducer() {
	go func() {
		newEvent := NewEvent("Timer", nil)
		ticker := time.NewTicker(e.TimeDuration)
		defer ticker.Stop()
	outer:
		for {
			select {
			case <-ticker.C:
				e.TimerEventChan <- newEvent
			case <-e.stopChan:
				close(e.TimerEventChan)
				if e.DebugLog {
					color.Redln(fmt.Sprintf("[%s]事件引擎定时发布生产者接收到关闭信号，定时器已终止，不再发布时间事件。", e.Name))
				}
				break outer
			}
		}
	}()
	if e.DebugLog {
		color.Greenln(fmt.Sprintf("[%s]事件引擎定时发布生产者已启动。", e.Name))
	}
}

// Put 发布事件，因为管道自带阻塞特性，为避免满后阻塞，因此没有消费者时不让发布
func (e *Engine) Put(event Event) {
	if e.ConsumerActive {
		e.queues[event.EventType].send(event)
	} else {
		if e.DebugLog {
			color.Redln(fmt.Sprintf("[%s]事件引擎处于关闭状态，丢弃事件发布。%+v", e.Name, event))
		}
	}
}
