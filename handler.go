package telebot

type IHandler interface {
	// CheckUpdate checks whether the update should handled by this handler.
	CheckUpdate(ctx Context) bool
	// HandleUpdate processes the update.
	HandleUpdate(ctx Context) error
	// Name gets the handler name; used to differentiate handlers programmatically. Names should be unique.
	Name() string
}

// HandlerFunc represents a handler function, which is
// used to handle actual endpoints.
type HandlerFunc func(Context) error

func (h HandlerFunc) Name() string {
	return "HandlerFunc"
}

func (h HandlerFunc) CheckUpdate(ctx Context) bool {
	return true
}

func (h HandlerFunc) HandleUpdate(ctx Context) error {
	return h(ctx)
}

// CheckHandlerList iterates over a list of handlers until a match is found; at which point it is returned.
func CheckHandlerList(handlers []IHandler, ctx Context) IHandler {
	for _, h := range handlers {
		if h.CheckUpdate(ctx) {
			return h
		}
	}
	return nil
}
