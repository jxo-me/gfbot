package middleware

import (
	"context"
	"encoding/json"
	tele "github.com/jxo-me/gfbot"
)

// Logger returns a middleware that logs incoming updates.
// If no custom logger provided, log.Default() will be used.
func Logger(ctx context.Context, logger ...tele.Logger) tele.MiddlewareFunc {
	var l tele.Logger
	if len(logger) > 0 {
		l = logger[0]
	} else {
		l = &tele.StdDebugLogger{}
	}

	return func(next tele.IHandler) tele.IHandler {
		return tele.HandlerFunc(func(c tele.IContext) error {
			data, _ := json.MarshalIndent(c.Update(), "", "  ")
			l.Debugf(ctx, string(data))
			return next.HandleUpdate(c)
		})
	}
}
