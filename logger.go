package telebot

import (
	"context"
	"fmt"
	"log"
)

// Logger is the interface to send logs to. It can be set using
// WithPublisherOptionsLogger() or WithConsumerOptionsLogger().
type Logger interface {
	Fatalf(context.Context, string, ...interface{})
	Errorf(context.Context, string, ...interface{})
	Warningf(context.Context, string, ...interface{})
	Infof(context.Context, string, ...interface{})
	Debugf(context.Context, string, ...interface{})
	Noticef(context.Context, string, ...interface{})
}

const loggingPrefix = "telegramBot"

// StdDebugLogger logs to stdout up to the `DebugF` level
type StdDebugLogger struct{}

func (l StdDebugLogger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s FATAL: %s", loggingPrefix, format), v...)
}

func (l StdDebugLogger) Errorf(ctx context.Context, format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s ERROR: %s", loggingPrefix, format), v...)
}

func (l StdDebugLogger) Warningf(ctx context.Context, format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s WARN: %s", loggingPrefix, format), v...)
}

func (l StdDebugLogger) Infof(ctx context.Context, format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s INFO: %s", loggingPrefix, format), v...)
}

func (l StdDebugLogger) Debugf(ctx context.Context, format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s DEBUG: %s", loggingPrefix, format), v...)
}

func (l StdDebugLogger) Noticef(ctx context.Context, format string, v ...interface{}) {}
