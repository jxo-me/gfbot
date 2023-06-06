package telebot

type IHandler interface {
	// CheckUpdate checks whether the update should handled by this handler.
	CheckUpdate(ctx IContext) bool
	// HandleUpdate processes the update.
	HandleUpdate(ctx IContext) error
	// Name gets the handler name; used to differentiate handlers programmatically. Names should be unique.
	Name() string
}

// HandlerFunc represents a handler function, which is
// used to handle actual endpoints.
type HandlerFunc func(IContext) error

func (h HandlerFunc) Name() string {
	return "HandlerFunc"
}

func (h HandlerFunc) CheckUpdate(ctx IContext) bool {
	return true
}

func (h HandlerFunc) HandleUpdate(ctx IContext) error {
	return h(ctx)
}
