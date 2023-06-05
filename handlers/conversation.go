package handlers

import (
	"errors"
	"fmt"

	tele "github.com/jxo-me/gfbot"
	"github.com/jxo-me/gfbot/handlers/conversation"
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
	// EntryPoints is the list of handlers to start the conversation.
	EntryPoints []tele.Handler
	// States is the map of possible states, with a list of possible handlers for each one.
	States map[string][]tele.Handler
	// StateStorage is responsible for storing all running conversations.
	StateStorage conversation.Storage

	// The following are all optional fields:
	// Exits is the list of handlers to exit the current conversation partway (eg /cancel commands)
	Exits []tele.Handler
	// Fallbacks is the list of handlers to handle updates which haven't been matched by any states.
	Fallbacks []tele.Handler
	// If True, a user can restart the conversation by hitting one of the entry points.
	AllowReEntry bool
}

type ConversationOpts struct {
	// Exits is the list of handlers to exit the current conversation partway (eg /cancel commands). This returns
	// EndConversation() by default, unless otherwise specified.
	Exits []tele.Handler
	// Fallbacks is the list of handlers to handle updates which haven't been matched by any other handlers.
	Fallbacks []tele.Handler
	// If True, a user can restart the conversation at any time by hitting one of the entry points again.
	AllowReEntry bool
	// StateStorage is responsible for storing all running conversations.
	StateStorage conversation.Storage
}

func NewConversation(entryPoints []tele.Handler, states map[string][]tele.Handler, opts *ConversationOpts) Conversation {
	c := Conversation{
		EntryPoints: entryPoints,
		States:      states,
		// Setup a default storage medium
		StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
	}

	if opts != nil {
		c.Exits = opts.Exits
		c.Fallbacks = opts.Fallbacks
		c.AllowReEntry = opts.AllowReEntry

		// If no StateStorage is specified, we should keep the default.
		if opts.StateStorage != nil {
			c.StateStorage = opts.StateStorage
		}
	}

	return c
}

func (c Conversation) CheckUpdate(ctx tele.Context) bool {
	// Note: Kinda sad that this error gets lost.
	h, _ := c.getNextHandler(ctx)
	fmt.Println("Conversation CheckUpdate 111111111:", h)
	return h != nil
}

func (c Conversation) HandleUpdate(ctx tele.Context) error {
	next, err := c.getNextHandler(ctx)
	fmt.Println("Conversation HandleUpdate err:", err)
	fmt.Println("Conversation HandleUpdate next:", next)
	if err != nil {
		return fmt.Errorf("failed to get next handler in conversation: %w", err)
	}
	if next == nil {
		// Note: this should be impossible
		return nil
	}

	var stateChange *ConversationStateChange
	err = next.HandleUpdate(ctx)
	fmt.Println("Conversation next.HandleUpdate:", err)
	if !errors.As(err, &stateChange) {
		// We don't wrap this error, as users might want to handle it explicitly
		return err
	}

	fmt.Println("2222222222222222222222222222222222222:", stateChange.End)
	if stateChange.End {
		// Mark the conversation as ended by deleting the conversation reference.
		err := c.StateStorage.Delete(ctx)
		if err != nil {
			return fmt.Errorf("failed to end conversation: %w", err)
		}
	}

	if stateChange.NextState != nil {
		fmt.Println("3333333333333333333333333333333:", c.States[*stateChange.NextState])
		// If the next state is defined, then move to it.
		if _, ok := c.States[*stateChange.NextState]; !ok {
			// Check if the "next" state is a supported state.
			return fmt.Errorf("unknown state: %w", stateChange)
		}
		err := c.StateStorage.Set(ctx, conversation.State{Key: *stateChange.NextState})
		fmt.Println("44444444444444444444444444444 err:", err)
		fmt.Println("44444444444444444444444444444 key:", *stateChange.NextState)
		if err != nil {
			return fmt.Errorf("failed to update conversation state: %w", err)
		}
	}

	if stateChange.ParentState != nil {
		// If a parent state is set, return that state for it to be handled.
		fmt.Println("5555555555555555:")
		return stateChange.ParentState
	}

	fmt.Println("666666666666666666666666666")
	return nil
}

// ConversationStateChange handles all the possible states that can be returned from a conversation.
type ConversationStateChange struct {
	// The next state to handle in the current conversation.
	NextState *string
	// End the current conversation
	End bool
	// Move the parent conversation (if any) to the desired state.
	ParentState *ConversationStateChange
}

func (s *ConversationStateChange) Error() string {
	// Avoid infinite print recursion by changing type
	type tmp *ConversationStateChange
	return fmt.Sprintf("conversation state change: %+v", tmp(s))
}

// NextConversationState moves to the defined state in the current conversation.
func NextConversationState(nextState string) *ConversationStateChange {
	return &ConversationStateChange{NextState: &nextState}
}

// NextParentConversationState moves to the defined state in the parent conversation, without changing the state of the current one.
func NextParentConversationState(parentState *ConversationStateChange) error {
	return &ConversationStateChange{ParentState: parentState}
}

// NextConversationStateAndParentState moves both the current conversation state and the parent conversation state.
// Can be helpful in the case of certain circular conversations.
func NextConversationStateAndParentState(nextState string, parentState *ConversationStateChange) error {
	return &ConversationStateChange{NextState: &nextState, ParentState: parentState}
}

// EndConversation ends the current conversation.
func EndConversation() error {
	return &ConversationStateChange{End: true}
}

// EndConversationToParentState ends the current conversation and moves the parent conversation to the new state.
func EndConversationToParentState(parentState *ConversationStateChange) error {
	return &ConversationStateChange{End: true, ParentState: parentState}
}

func (c Conversation) Name() string {
	return fmt.Sprintf("conversation_%p", c.States)
}

// getNextHandler goes through all the handlers in the conversation, until it finds a handler that matches.
// If no matching handler is found, returns nil.
func (c Conversation) getNextHandler(ctx tele.Context) (tele.Handler, error) {
	// Check if a conversation has already started for this user.
	currState, err := c.StateStorage.Get(ctx)
	fmt.Println("getNextHandler err:", err)
	if err != nil {
		if errors.Is(err, conversation.KeyNotFound) {
			fmt.Println("getNextHandler EntryPoints:", c.EntryPoints)
			fmt.Println("getNextHandler checkHandlerList result:", checkHandlerList(c.EntryPoints, ctx))
			// If this is an unknown conversation key, then we know this is a new conversation, so we check all
			// entrypoints.
			return checkHandlerList(c.EntryPoints, ctx), nil
		}
		// Else, we need to handle the error.
		return nil, fmt.Errorf("failed to get state from conversation storage: %w", err)
	}

	fmt.Println("getNextHandler currState:", currState)
	// If reentry is allowed, check the entrypoints again.
	if c.AllowReEntry {
		if next := checkHandlerList(c.EntryPoints, ctx); next != nil {
			return next, nil
		}
	}

	// Else, exits -> handle any conversation exits/cancellations.
	if next := checkHandlerList(c.Exits, ctx); next != nil {
		return wrappedExitHandler{h: next}, nil
	}

	// Else, check state mappings (the magic happens here!).
	if next := checkHandlerList(c.States[currState.Key], ctx); next != nil {
		return next, nil
	}

	// Else, fallbacks -> handle any updates which haven't been caught by the state or exit handlers.
	if next := checkHandlerList(c.Fallbacks, ctx); next != nil {
		return next, nil
	}

	return nil, nil
}

// checkHandlerList iterates over a list of handlers until a match is found; at which point it is returned.
func checkHandlerList(handlers []tele.Handler, ctx tele.Context) tele.Handler {
	for _, h := range handlers {
		if h.CheckUpdate(ctx) {
			return h
		}
	}
	return nil
}

// wrappedExitHandler ensures that exit handlers return conversation ends by default.
type wrappedExitHandler struct {
	h tele.Handler
}

func (w wrappedExitHandler) CheckUpdate(ctx tele.Context) bool {
	fmt.Println("wrappedExitHandler CheckUpdate", w.h.CheckUpdate(ctx))
	return w.h.CheckUpdate(ctx)
}

func (w wrappedExitHandler) HandleUpdate(ctx tele.Context) error {
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
