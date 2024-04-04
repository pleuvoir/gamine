package socketio

import (
	"fmt"
	"github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/pleuvoir/gamine/helper/helper_lang"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"net"
	"net/http"
	"time"
)

var instance *Instance

// ServerState 服务状态
type ServerState int

const (
	methodRequest = "request"
	methodPush    = "push"
	// ServerStarting 服务启动中
	ServerStarting ServerState = 1
	// ServerStarted 服务已启动
	ServerStarted ServerState = 2
	// ServerFailed 服务启动失败
	ServerFailed ServerState = 3
)

type Instance struct {
	socketIO    *gosocketio.Server
	gin         *gin.Engine
	port        int `yaml:"port"`
	startedChan chan bool
	//延时测试端口的时间
	testPortDelayed time.Duration
	//测试端口的重试次数，若设置为小于1的数则按1次处理
	testPortRetryTimes int
	// 服务状态
	state ServerState
}

func New(port int) (goSocketIO *Instance) {
	g := gin.New()
	g.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"*"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       false,
		ExposeHeaders:          nil,
		MaxAge:                 12 * time.Hour,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))
	socketIO := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	g.Any("/socket.io/*any", gin.WrapH(socketIO))
	_ = socketIO.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		color.Greenln("socketIO已连接", c.Ip())
	})
	_ = socketIO.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		color.Yellowln("socketIO已断开连接", c.Ip())
	})
	return &Instance{
		socketIO:           socketIO,
		gin:                g,
		startedChan:        make(chan bool, 1),
		testPortDelayed:    time.Second * 2,
		testPortRetryTimes: 3,
		port:               port}
}

func (i *Instance) Run() error {
	i.state = ServerStarting
	serv := &http.Server{Addr: ":" + helper_lang.IntToString(i.port), Handler: i.gin}
	i.startServer(serv) //异步启动
	i.listenServerStarted()
	instance = i
	return nil
}

func (i *Instance) startServer(serv *http.Server) {
	go func() {
		i.startedChan <- true
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			i.state = ServerFailed
			color.Redln(fmt.Sprintf("socketIO启动失败： %s", err))
		}
	}()
}

// listenServerStarted 启动监听，在服务启动时检测http服务监听的端口
func (i *Instance) listenServerStarted() {
	go func() {
		<-i.startedChan
		if i.testPortRetry() {
			i.state = ServerStarted
			color.Greenln(fmt.Sprintf("socketIO已启动 @%s", helper_lang.IntToString(i.port)))
		}
	}()
}

// testPortRetry 检测http服务监听的端口，该方法会延时阻塞执行，若监测端口超时则会重试
func (i *Instance) testPortRetry() bool {
	time.Sleep(i.testPortDelayed)
	testPortTimes := helper_lang.If(i.testPortRetryTimes < 1, 1, i.testPortRetryTimes).(int)
	for c := 0; c < testPortTimes; c++ {
		if i.testPort() {
			return true
		}
	}
	return false
}

// testPort 检测一次http服务监听的端口
func (i *Instance) testPort() bool {
	conn, err := net.DialTimeout("tcp", ":"+helper_lang.IntToString(i.port), time.Millisecond*500)
	defer helper_os.CloseQuietly(conn)
	return err == nil
}

type RequestMessage struct {
	MethodName string `json:"methodName"`
	Data       any    `json:"data"`
}

type ResponseMessage struct {
	MethodName string `json:"methodName"`
	Data       any    `json:"data"`
}

type RequestHandler func(msg RequestMessage) ResponseMessage

func (i *Instance) WithRequest(handler RequestHandler) error {
	err := i.socketIO.On(methodRequest, func(c *gosocketio.Channel, msg RequestMessage) ResponseMessage {
		return handler(msg)
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) PushResponse(message ResponseMessage) {
	i.socketIO.BroadcastToAll(methodPush, message)
}

func (i *Instance) GetName() string {
	return "socket-io"
}

func Get() (goSocketIO *Instance) {
	return instance
}
