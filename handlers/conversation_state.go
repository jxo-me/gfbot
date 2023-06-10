package handlers

import "fmt"

// ConversationStateChange handles all the possible states that can be returned from a conversation.
type ConversationStateChange struct {
	// The next state to handle in the current conversation.
	NextState *string
	// End the current conversation
	End bool
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

// EndConversation ends the current conversation.
func EndConversation() error {
	return &ConversationStateChange{End: true}
}
