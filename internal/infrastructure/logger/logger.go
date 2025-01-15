package logger

import "go.uber.org/zap"

var logger *zap.SugaredLogger

type Fields map[string]interface{}

func NewZapLogger() {
	log, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	sugar := log.Sugar()
	defer log.Sync()

	logger = sugar
}

func Info(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

func WithFields(fields Fields) *zap.SugaredLogger {
	var f = make([]interface{}, 0)
	for index, field := range fields {
		f = append(f, index)
		f = append(f, field)
	}
	log := logger.With(f...)
	return log
}
