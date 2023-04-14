package log

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLog 根据配置创建一个新的zap
func NewZapLog(name string) Logger {
	return &zapLog{
		zap.New(
			zapcore.NewCore(getEncoder(), getLogWriter(name), zapcore.DebugLevel),
			zap.AddCallerSkip(2),
			zap.AddCaller(),
		),
	}
}

func getLogWriter(name string) zapcore.WriteSyncer {
	logPath := "/usr/local/service/log"
	if runtime.GOOS != "linux" {
		dir, _ := os.Getwd()
		logPath = filepath.Join(dir, "log")
	}
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   filepath.Join(logPath, name),
			MaxSize:    10,
			MaxAge:     7,
			MaxBackups: 10,
			Compress:   false,
		}),
		zapcore.Lock(os.Stdout),
	)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 格式化时间 可自定义
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

type zapLog struct {
	*zap.Logger
}

func (z *zapLog) Debug(args ...interface{}) {
	if z.Logger.Core().Enabled(zapcore.DebugLevel) {
		z.Logger.Sugar().Debug(args...)
	}
}

func (z *zapLog) Debugf(format string, args ...interface{}) {
	if z.Logger.Core().Enabled(zapcore.DebugLevel) {
		z.Logger.Sugar().Debugf(format, args...)
	}
}

func (z *zapLog) Info(args ...interface{}) {
	if z.Logger.Core().Enabled(zapcore.InfoLevel) {
		z.Logger.Sugar().Info(args...)
	}
}

func (z *zapLog) Infof(format string, args ...interface{}) {
	if z.Logger.Core().Enabled(zapcore.InfoLevel) {
		z.Logger.Sugar().Infof(format, args...)
	}
}

func (z *zapLog) Error(args ...interface{}) {
	if z.Logger.Core().Enabled(zapcore.ErrorLevel) {
		z.Logger.Sugar().Error(args...)
	}
}

func (z *zapLog) Errorf(format string, args ...interface{}) {
	if z.Logger.Core().Enabled(zapcore.ErrorLevel) {
		z.Logger.Sugar().Errorf(format, args...)
	}
}

func (z *zapLog) Fatal(args ...interface{}) {
	if z.Logger.Core().Enabled(zapcore.FatalLevel) {
		z.Logger.Sugar().Fatal(args...)
	}
}

func (z *zapLog) Fatalf(format string, args ...interface{}) {
	if z.Logger.Core().Enabled(zapcore.FatalLevel) {
		z.Logger.Sugar().Fatalf(format, args...)
	}
}
