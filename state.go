package telebot

// State stores all the variables relevant to the current conversation state.
//
// Note: More keys may be added in the future to support additional features.
// As such, any storage implementations should be flexible, and allow for storing the entire struct rather than
// individual fields.
type State struct {
	// Key represents the name of the current state, as defined in the SubHandlers map of handlers.Conversation.
	Key         string
	ServiceName string
	Data        map[string]any
}

func (s *State) SetKey(key string) *State {
	s.Key = key
	return s
}

func (s *State) SetServiceName(srv string) *State {
	s.ServiceName = srv
	return s
}

func (s *State) SetData(data map[string]any) *State {
	s.Data = data
	return s
}

func (s *State) UpdateData(action string, data any) *State {
	if s.Data == nil {
		s.Data = map[string]any{}
	}
	s.Data[action] = data

	return s
}
