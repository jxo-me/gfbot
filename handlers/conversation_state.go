package handlers

import "fmt"

// StateChange handles all the possible states that can be returned from a conversation.
type StateChange struct {
	// The next state to handle in the current conversation.
	NextState *string
	// End the current conversation
	End bool
}

func (s *StateChange) Error() string {
	// Avoid infinite print recursion by changing type
	type tmp *StateChange
	return fmt.Sprintf("conversation state change: %+v", tmp(s))
}

// NextConversationState moves to the defined state in the current conversation.
func NextConversationState(nextState string) *StateChange {
	return &StateChange{NextState: &nextState}
}

// EndConversation ends the current conversation.
func EndConversation() error {
	return &StateChange{End: true}
}
