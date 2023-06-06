package handlers

import "fmt"

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
