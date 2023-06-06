package handlers

import (
	"fmt"
	tele "github.com/jxo-me/gfbot"
)

// wrappedExitHandler ensures that exit handlers return conversation ends by default.
type wrappedExitHandler struct {
	h tele.IHandler
}

func (w wrappedExitHandler) CheckUpdate(ctx tele.IContext) bool {
	fmt.Println("wrappedExitHandler CheckUpdate", w.h.CheckUpdate(ctx))
	return w.h.CheckUpdate(ctx)
}

func (w wrappedExitHandler) HandleUpdate(ctx tele.IContext) error {
	err := w.h.HandleUpdate(ctx)
	fmt.Println("wrappedExitHandler HandleUpdate", err)
	if err != nil {
		return err
	}
	return EndConversation()
}

func (w wrappedExitHandler) Name() string {
	return w.h.Name()
}
