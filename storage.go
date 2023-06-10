package telebot

// IStorage allows you to define custom backends for retaining conversation conversations.
// If you are looking to persist conversation data, you should implement this interface with you backend of choice.
// Note: Make sure to store the entire State struct; future changes may add new fields.
type IStorage interface {
	// Get returns the state for the specified conversation key.
	// Note that this is checked at each incoming message, so may be a bottleneck for some implementations.
	//
	// If the key is not found (and as such, this conversation has not yet started), this method should return the
	// ConversationKeyNotFound error.
	// Get(key string) (*State, error)
	Get(ctx IContext) (*State, error)

	// Set add the conversation state.
	Set(ctx IContext, state State) error

	// Next updates the conversation state.
	Next(ctx IContext, key string) error

	// UpdateData updates the conversation action state.
	UpdateData(ctx IContext, act string, data any) error

	// Delete ends the conversation, removing the key from the storage.
	Delete(ctx IContext) error
}
