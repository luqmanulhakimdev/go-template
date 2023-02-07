package logger

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	X_CORRELATION_ID = "X-Correlation-ID"
	CORRELATION_ID   = "Correlation-ID"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func Init(level, format string) {
	// SetFormatter
	log.SetFormatter(GetFormatter(format))

	// SetLevel
	switch strings.ToLower(level) {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.SetOutput(os.Stdout)
}

func GetFormatter(logFormat string) log.Formatter {
	switch strings.ToLower(logFormat) {
	case "json":
		return &log.JSONFormatter{}
	case "gce":
		return NewGCEFormatter()
	default:
		return &log.TextFormatter{
			FullTimestamp: true,
		}
	}
}

func GetLogger() *log.Logger {
	return log.StandardLogger()
}

func Info(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value(X_CORRELATION_ID)
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		CORRELATION_ID: id,
	}).Infof(format, values...)
}

func Warn(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value(X_CORRELATION_ID)
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		CORRELATION_ID: id,
	}).Warnf(format, values...)
}

func Error(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value(X_CORRELATION_ID)
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		CORRELATION_ID: id,
	}).Errorf(format, values...)
}

func ErrorWithStack(ctx context.Context, err error, format string, values ...interface{}) {
	var id string
	val := ctx.Value(X_CORRELATION_ID)
	if val != nil {
		id = val.(string)
	}
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])
	log.WithFields(log.Fields{
		CORRELATION_ID: id,
		"stack":        stFormat,
		"stat_msg":     err.Error(),
	}).Errorf(format, values...)
}

func Debug(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value(X_CORRELATION_ID)
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		CORRELATION_ID: id,
	}).Debugf(format, values...)
}

func Fatal(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value(X_CORRELATION_ID)
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		CORRELATION_ID: id,
	}).Fatalf(format, values...)
}
