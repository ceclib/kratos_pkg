package zap

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	log *zap.Logger
}

// LogConfig 日志的初始化配置
type LogConfig struct {
	//日志路径
	LogPath string
	//项目名字
	AppName string
	//是否debug
	Debug bool
	//CallerSkip
	CallerSkip int
}

// NewLogger 返回一个日志实例
func NewLogger(LogConfig LogConfig) *Logger {
	return &Logger{initLog(LogConfig)}
}

// 初始化日志
func initLog(LogConfig LogConfig) *zap.Logger {
	// 根据配置类引入
	logPath := LogConfig.LogPath
	name := LogConfig.AppName
	debug := LogConfig.Debug

	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    100,     // 每个日志文件保存的大小 单位:M
		MaxAge:     7,       // 文件最多保存多少天
		MaxBackups: 30,      // 日志文件最多保存多少个备份
		Compress:   false,   // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		//TimeKey:        "time",
		//NameKey:        "logger",
		//CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	// 如果是开发环境，同时在控制台上也输出
	if debug {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCallerSkip(LogConfig.CallerSkip)
	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	field := zap.Fields(zap.String("appName", name))

	return zap.New(core, caller, development, field, zap.AddCaller())
}

func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}
	//fmt.Printf("%+v",data)
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	case log.LevelFatal:
		l.log.Fatal("", data...)
	}
	return nil
}

func (l *Logger) Sync() error {
	return l.log.Sync()
}
