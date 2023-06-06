package handlers

import tele "github.com/jxo-me/gfbot"

// checkHandlerList iterates over a list of handlers until a match is found; at which point it is returned.
func checkHandlerList(handlers []tele.IHandler, ctx tele.IContext) tele.IHandler {
	for _, h := range handlers {
		if h.CheckUpdate(ctx) {
			return h
		}
	}
	return nil
}
