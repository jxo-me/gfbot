package handlers

import (
	"errors"
	"fmt"

	tele "github.com/jxo-me/gfbot"
)

// TODO: Add a "block" option to force linear processing. Also a "waiting" state to handle blocked handlers.
// TODO: Allow for timeouts (and a "timeout" state to handle that)

// The Conversation handler is an advanced handler which allows for running a sequence of commands in a stateful manner.
// An example of this flow can be found at t.me/Botfather; upon receiving the "/newbot" command, the user is asked for
// the name of their bot, which is sent as a separate message.
//
// The bot's internal state allows it to check at which point of the conversation the user is, and decide how to handle
// the next update.
type Conversation struct {
	EntryName string
	// EntryHandler is the list of handlers to start the conversation.
	EntryHandler tele.IHandler
	// States is the map of possible states, with a list of possible handlers for each one.
	States map[string][]tele.IHandler

	ExitName string
	// The following are all optional fields:
	// ExitHandler is the list of handlers to exit the current conversation partway (eg /cancel commands)
	ExitHandler tele.IHandler
	// Fallbacks is the list of handlers to handle updates which haven't been matched by any states.
	Fallbacks tele.IHandler
	// If True, a user can restart the conversation by hitting one of the entry points.
	AllowReEntry bool
}

type ConversationOpts struct {
	ExitName string
	// ExitHandler is the list of handlers to exit the current conversation partway (eg /cancel commands). This returns
	// EndConversation() by default, unless otherwise specified.
	ExitHandler tele.IHandler
	// Fallbacks is the list of handlers to handle updates which haven't been matched by any other handlers.
	Fallbacks tele.IHandler
	// If True, a user can restart the conversation at any time by hitting one of the entry points again.
	AllowReEntry bool
}

func NewConversation(entryName string, entryPoint tele.IHandler, states map[string][]tele.IHandler, opts *ConversationOpts) Conversation {
	c := Conversation{
		EntryName:    entryName,
		EntryHandler: entryPoint,
		States:       states,
	}

	if opts != nil {
		c.ExitName = opts.ExitName
		c.ExitHandler = opts.ExitHandler
		c.Fallbacks = opts.Fallbacks
		c.AllowReEntry = opts.AllowReEntry
	}

	return c
}

func (c Conversation) CheckUpdate(ctx tele.IContext) bool {
	// Note: Kinda sad that this error gets lost.
	h, _ := c.getNextHandler(ctx)
	return h != nil
}

func (c Conversation) HandleUpdate(ctx tele.IContext) error {
	next, err := c.getNextHandler(ctx)
	if err != nil {
		return fmt.Errorf("failed to get next handler in conversation: %w", err)
	}
	if next == nil {
		// Note: this should be impossible
		return nil
	}

	var stateChange *ConversationStateChange
	err = next.HandleUpdate(ctx)
	if !errors.As(err, &stateChange) {
		// We don't wrap this error, as users might want to handle it explicitly
		return err
	}

	if stateChange.End {
		// Mark the conversation as ended by deleting the conversation reference.
		err := ctx.Bot().StateStorage.Delete(ctx)
		if err != nil {
			return fmt.Errorf("failed to end conversation: %w", err)
		}
	}

	if stateChange.NextState != nil {
		// If the next state is defined, then move to it.
		if _, ok := c.States[*stateChange.NextState]; !ok {
			// Check if the "next" state is a supported state.
			return fmt.Errorf("unknown state: %w", stateChange)
		}
		err := ctx.Bot().StateStorage.Set(ctx, tele.State{Key: *stateChange.NextState, EntryName: c.EntryName})
		if err != nil {
			return fmt.Errorf("failed to update conversation state: %w", err)
		}
	}

	if stateChange.ParentState != nil {
		// If a parent state is set, return that state for it to be handled.
		return stateChange.ParentState
	}

	return nil
}

func (c Conversation) Name() string {
	return c.EntryName
}

// getNextHandler goes through all the handlers in the conversation, until it finds a handler that matches.
// If no matching handler is found, returns nil.
func (c Conversation) getNextHandler(ctx tele.IContext) (tele.IHandler, error) {
	// Check if a conversation has already started for this user.
	currState, _ := ctx.Bot().StateStorage.Get(ctx)
	cmd := ctx.Message().Text
	switch cmd {
	case c.EntryName:
		return c.EntryHandler, nil
	case c.ExitName:
		if currState != nil {
			return c.ExitHandler, nil
		}
	default:
		if currState != nil {
			if next := tele.CheckHandlerList(c.States[currState.Key], ctx); next != nil {
				return next, nil
			}
		}
	}

	return nil, nil
}
