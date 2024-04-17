package handlers

import (
	tele "github.com/jxo-me/gfbot"
)

// wrappedExitHandler ensures that exit handlers return conversation ends by default.
type wrappedExitHandler struct {
	h tele.IHandler
}

func (w wrappedExitHandler) CheckUpdate(ctx tele.Context) bool {
	return w.h.CheckUpdate(ctx)
}

func (w wrappedExitHandler) HandleUpdate(ctx tele.Context) error {
	err := w.h.HandleUpdate(ctx)
	if err != nil {
		return err
	}
	return EndConversation()
}

func (w wrappedExitHandler) Name() string {
	return w.h.Name()
}
