package logger

import (
	"context"
	"errors"

	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	Init("info", "json")
	Init("warn", "json")
	Init("error", "json")
	Init("debug", "json")
	Init("", "json")
}

func TestGetFormatter(t *testing.T) {
	res := GetFormatter("json")
	assert.Equal(t, res, &log.JSONFormatter{})

	res = GetFormatter("gce")
	assert.Equal(t, res, &GCEFormatter{})

	res = GetFormatter("")
	assert.Equal(t, res, &log.TextFormatter{FullTimestamp: true})
}

func TestGetLogger(t *testing.T) {
	res := GetLogger()
	assert.Equal(t, res, log.StandardLogger())
}

func TestInfo(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "123")
	Info(ctx, "Log Test: %v", "Test Func Info")
}

func TestWarn(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "123")
	Warn(ctx, "Log Test: %v", "Test Func Warn")
}

func TestError(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "123")
	Error(ctx, "Log Test: %v", "Test Func Error")
}

func TestErrorWithStack(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "123")
	ErrorWithStack(ctx, errors.New("error_test"), "Log Test: %v", "Test Func ErrorWithStack")
}

func TestDebug(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "123")
	Debug(ctx, "Log Test: %v", "Test Func Debug")
}

func TestFatal(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "123")

	defer func() { log.StandardLogger().ExitFunc = nil }()
	var fatal bool
	log.StandardLogger().ExitFunc = func(int) { fatal = true }

	Fatal(ctx, "Log Test: %v", "Test Func Fatal")
	assert.Equal(t, fatal, true)
}
