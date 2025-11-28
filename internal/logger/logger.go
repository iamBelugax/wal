package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func New(logFile string, lvl zapcore.Level) (*Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		TimeKey:        "timestamp",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	fileWS := zapcore.AddSync(f)

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleWS := zapcore.AddSync(os.Stdout)

	atomicLevel := zap.NewAtomicLevelAt(lvl)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWS, atomicLevel),
		zapcore.NewCore(consoleEncoder, consoleWS, atomicLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return &Logger{SugaredLogger: logger.Sugar()}, nil
}

func (l *Logger) Close() error {
	return l.Sync()
}
