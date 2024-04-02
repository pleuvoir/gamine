package restful

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/pleuvoir/gamine/helper/helper_lang"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const RestEngineName = "Gin"

// ServerState 服务状态
type ServerState int

const (
	// ServerStarting 服务启动中
	ServerStarting ServerState = 1
	// ServerStarted 服务已启动
	ServerStarted ServerState = 2
	// ServerFailed 服务启动失败
	ServerFailed ServerState = 3
)

// ServerStartedListener 服务启动后调用的监听程序
type ServerStartedListener func(engine *Instance)

type Instance struct {
	Gin         *gin.Engine
	httpServer  *http.Server
	startedChan chan bool
	//延时测试端口的时间
	testPortDelayed time.Duration
	//测试端口的重试次数，若设置为小于1的数则按1次处理
	testPortRetryTimes int
	//服务启动后调用的监听程序
	serverStartedListener ServerStartedListener
	// 端口号
	port string
	// 服务状态
	State ServerState
}

func NewRestServer(port string) *Instance {
	r := &Instance{
		startedChan:        make(chan bool, 1),
		testPortDelayed:    time.Second * 2,
		testPortRetryTimes: 3,
		port:               port,
	}
	gin.SetMode(gin.ReleaseMode)
	r.Gin = gin.New()
	r.Gin.Use(gin.Recovery())
	return r
}

func (e *Instance) WithServerStartedListener(listener ServerStartedListener) {
	e.serverStartedListener = listener
}

// WithUseRequestLog 使用的日志格式 一般需要放置在第一位置 放置不能c.Next()
func (e *Instance) WithUseRequestLog(log *logrus.Logger) {
	e.Gin.Use(func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		log.Infoln(fmt.Sprintf("| %3d | %13v | %15s | %s | %s", statusCode, latencyTime, clientIP, reqMethod, reqUri))
	})
}

// WithCors 支持跨域所有格式
func (e *Instance) WithCors() {
	e.Gin.Use(cors.New(cors.Config{
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
}

func (e *Instance) GetName() string {
	return RestEngineName
}

func (e *Instance) Run() error {
	e.State = ServerStarting
	serv := &http.Server{Addr: ":" + e.port, Handler: e.Gin}
	e.httpServer = serv
	e.startServer(serv) //异步启动
	e.listenServerStarted()
	e.gracefulShutdown(serv) //阻塞进程等待退出
	return nil
}

// WithGinConfig 对外暴露GIN
func (e *Instance) WithGinConfig(with func(e *gin.Engine)) {
	with(e.Gin)
}

func (e *Instance) startServer(serv *http.Server) {
	go func() {
		e.startedChan <- true
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			e.State = ServerFailed
			color.Redln(fmt.Sprintf("restful服务启动错误： %s", err))
		}
	}()
}

// listenServerStarted 启动监听，在服务启动时检测http服务监听的端口
func (e *Instance) listenServerStarted() {
	if e.serverStartedListener == nil {
		return
	}
	go func() {
		<-e.startedChan
		if e.testPortRetry() {
			e.State = ServerStarted
			color.Greenln(fmt.Sprintf("restful服务已启动 @%s", e.port))
			e.serverStartedListener(e)
		}
	}()
}

// testPortRetry 检测http服务监听的端口，该方法会延时阻塞执行，若监测端口超时则会重试
func (e *Instance) testPortRetry() bool {
	time.Sleep(e.testPortDelayed)
	testPortTimes := helper_lang.If(e.testPortRetryTimes < 1, 1, e.testPortRetryTimes).(int)
	for i := 0; i < testPortTimes; i++ {
		if e.testPort() {
			return true
		}
	}
	return false
}

// testPort 检测一次http服务监听的端口
func (e *Instance) testPort() bool {
	conn, err := net.DialTimeout("tcp", ":"+e.port, time.Millisecond*500)
	defer helper_os.CloseQuietly(conn)
	return err == nil
}

// gracefulShutdown 优雅关闭服务
func (e *Instance) gracefulShutdown(serv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	color.Greenln("restful服务接收到优雅退出信号")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := serv.Shutdown(ctx); err != nil {
		color.Redln(fmt.Sprintf("restful服务优雅退出错误： %s", err))
	}
	color.Greenln("restful服务优雅退出")
}

func (e *Instance) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := e.httpServer.Shutdown(ctx); err != nil {
		color.Redln(fmt.Sprintf("restful服务停止错误： %s", err))
	}
	color.Greenln("restful服务停止")
}
