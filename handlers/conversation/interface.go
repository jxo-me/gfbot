package conversation

import (
	tele "github.com/jxo-me/gfbot"
)

// Storage allows you to define custom backends for retaining conversation conversations.
// If you are looking to persist conversation data, you should implement this interface with you backend of choice.
// Note: Make sure to store the entire State struct; future changes may add new fields.
type Storage interface {
	// Get returns the state for the specified conversation key.
	// Note that this is checked at each incoming message, so may be a bottleneck for some implementations.
	//
	// If the key is not found (and as such, this conversation has not yet started), this method should return the
	// ConversationKeyNotFound error.
	// Get(key string) (*State, error)
	Get(ctx tele.Context) (*State, error)

	// Set updates the conversation state.
	Set(ctx tele.Context, state State) error

	// Delete ends the conversation, removing the key from the storage.
	Delete(ctx tele.Context) error
}
