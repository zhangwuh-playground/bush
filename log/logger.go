package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	NT string = "nt"
)

var defaultLogger *zap.Logger

type Logger *zap.Logger

func init() {
	defaultLogger = NewLogger("Default-Logger")
}

func NewLoggerWithLevel(loggerName string, logLevel zapcore.Level) Logger {
	zapLogger, err := newZapLogger(logLevel)

	if err != nil {
		panic(fmt.Sprintf("Failed to create logger: {%s}. Error: %v", loggerName, err))
	}

	zapLogger.Info(fmt.Sprintf("Logger[%s] level: %s", loggerName, logLevel.String()))

	return zapLogger
}

func NewLogger(loggerName string) Logger {
	return NewLoggerWithLevel(loggerName, zap.InfoLevel)
}

func newZapLogger(level zapcore.Level) (*zap.Logger, error) {
	return newLogstashConfig(level).Build(zap.AddCallerSkip(2))
}

func newLogstashConfig(level zapcore.Level) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    newLogstashEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.999+08:00"))
}

func newLogstashEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "thread_name",
		MessageKey:     "message",
		StacktraceKey:  "stack_trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func Message(format string, fields ...interface{}) string {
	return fmt.Sprintf(format, fields...)
}

func traceId(traceId string) zap.Field {
	return zap.String("tracingID", traceId)
}

func Sync() error {
	return defaultLogger.Sync()
}

func isValidTraceId(ti string) bool {
	return ti != NT && len(ti) > 0
}

func Debug(message string, fields ...zap.Field) {
	defaultLogger.Debug(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	defaultLogger.Fatal(message, fields...)
}

func Info(ti, message string, fields ...zap.Field) {
	if isValidTraceId(ti) {
		fields = append(fields, traceId(ti))
	}
	defaultLogger.Info(message, fields...)
}

func Warn(ti, message string, fields ...zap.Field) {
	if isValidTraceId(ti) {
		fields = append(fields, traceId(ti))
	}
	defaultLogger.Warn(message, fields...)
}

func Error(ti, message string, err error, fields ...zap.Field) {
	if isValidTraceId(ti) {
		fields = append(fields, traceId(ti))
	}
	fields = append(fields, zap.String("Error", err.Error()))
	defaultLogger.Error(message, fields...)
}

func InfoNt(message string, fields ...zap.Field) {
	Info(NT, message, fields...)
}

func WarnNt(message string, fields ...zap.Field) {
	Warn(NT, message, fields...)
}

func ErrorNt(message string, err error, fields ...zap.Field) {
	Error(NT, message, err, fields...)
}
