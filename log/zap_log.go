package log

import (
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLog 根据配置创建一个新的zap
func NewZapLog(cfg *LoggerConfig) Logger {
	callerSkip := 2
	if cfg.CallerSkip > 0 {
		callerSkip = cfg.CallerSkip
	}
	return &zapLog{
		zap.New(
			zapcore.NewCore(newEncoder(), newLogWriter(&cfg.WriteConfig), zapcore.DebugLevel),
			zap.AddCallerSkip(callerSkip),
			zap.AddCaller(),
		),
	}
}

func newLogWriter(cfg *WriteConfig) zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   filepath.Join(cfg.LogPath, cfg.Filename),
			MaxSize:    cfg.MaxSize,
			MaxAge:     cfg.MaxAge,
			MaxBackups: cfg.MaxBackups,
			Compress:   cfg.Compress,
		}),
		zapcore.Lock(os.Stdout),
	)
}

func newEncoder() zapcore.Encoder {
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
