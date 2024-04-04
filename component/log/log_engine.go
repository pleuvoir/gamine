package log

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/pleuvoir/gamine/helper/helper_os"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

func (e *Engine) Print(args ...interface{}) {
	e.log.Print(args...)
}

func (e *Engine) Printf(format string, args ...interface{}) {
	e.log.Printf(format, args...)
}

func (e *Engine) Println(args ...interface{}) {
	e.log.Println(args...)
}

func (e *Engine) Debug(args ...interface{}) {
	e.log.Debug(args...)
}

func (e *Engine) Debugf(format string, args ...interface{}) {
	e.log.Debugf(format, args...)
}

func (e *Engine) Debugln(args ...interface{}) {
	e.log.Debugln(args...)
}

func (e *Engine) Info(args ...interface{}) {
	e.log.Info(args...)
}

func (e *Engine) Infof(format string, args ...interface{}) {
	e.log.Infof(format, args...)
}

func (e *Engine) Infoln(args ...interface{}) {
	e.log.Infoln(args...)
}

func (e *Engine) Warn(args ...interface{}) {
	e.log.Warn(args...)
}

func (e *Engine) Warnf(format string, args ...interface{}) {
	e.log.Warnf(format, args...)
}

func (e *Engine) Warnln(args ...interface{}) {
	e.log.Warnln(args...)
}

func (e *Engine) Error(args ...interface{}) {
	e.log.Error(args...)
}

func (e *Engine) Errorf(format string, args ...interface{}) {
	e.log.Errorf(format, args...)
}

func (e *Engine) Errorln(args ...interface{}) {
	e.log.Errorln(args...)
}

func (e *Engine) Fatal(args ...interface{}) {
	e.log.Fatal(args...)
}

func (e *Engine) Fatalf(format string, args ...interface{}) {
	e.log.Fatalf(format, args...)
}

func (e *Engine) Fatalln(args ...interface{}) {
	e.log.Fatalln(args...)
}

func (e *Engine) Panic(args ...interface{}) {
	e.log.Panic(args...)
}

func (e *Engine) Panicf(format string, args ...interface{}) {
	e.log.Panicf(format, args...)
}

func (e *Engine) Panicln(args ...interface{}) {
	e.log.Panicln(args...)
}

type Engine struct {
	log *logrus.Logger
}

// NewEngine 初始化日志，初始化失败时会抛出异常
//
//	config: 系统配置
func NewEngine(config *Config) *Engine {

	engine := Engine{}

	var maxAgeDuration time.Duration
	var rotationTimeDuration time.Duration
	maxAge := config.MaxAge
	if maxAge == "" {
		maxAgeDuration = time.Hour * 24 * 60 //60天
	} else {
		duration, err := time.ParseDuration(maxAge)
		if err != nil {
			panic(fmt.Sprintf("init log maxAge fail: %s, %s", maxAge, err.Error()))
		}
		maxAgeDuration = duration
	}
	rotationTime := config.RotationTime
	if rotationTime == "" {
		rotationTimeDuration = time.Hour * 24 //1天
	} else {
		duration, err := time.ParseDuration(rotationTime)
		if err != nil {
			panic(fmt.Sprintf("init log rotationTime fail: %s, %s", rotationTime, err.Error()))
		}
		rotationTimeDuration = duration
	}
	engine.configLocalFilesystemLogger(config.Level,
		config.Path,
		config.Filename,
		maxAgeDuration,
		rotationTimeDuration)

	return &engine
}

// config loggers log to local filesystem, with file rotation
func (e *Engine) configLocalFilesystemLogger(level string, logPath string, logFileName string,
	maxAge time.Duration, rotationTime time.Duration) {
	color.Printf("<light_green>准备初始化日志:</> level:%s, logPath:%s, logFilename:%s, maxAge:%s, rotationTime:%s \n",
		level, logPath, logFileName, maxAge.String(), rotationTime.String())

	if !helper_os.FileExists(logPath) {
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	e.log = logrus.New()

	lv, err := logrus.ParseLevel(level)
	if err != nil {
		log.Panicf("cannot parse log level: %s, %+v", level, errors.WithStack(err))
	}
	e.log.SetLevel(lv)

	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Panicf("config local file system logger error. %+v", errors.WithStack(err))
	}

	e.log.SetReportCaller(false)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: writer,
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &ResetFormatter{})
	e.log.AddHook(lfHook)
}

type ResetFormatter struct{}

func (m *ResetFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d %s] %s\n",
			timestamp, entry.Level, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}
	b.WriteString(newLog)
	return b.Bytes(), nil
}
