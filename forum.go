package telebot

// ForumTopic This object represents a forum topic.
type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicCreated This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	// Name of the topic
	Name string `json:"name"`
	// Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`
	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicClosed This object represents a service message about a forum topic closed in the chat. Currently holds no information.
type ForumTopicClosed struct {
}

// ForumTopicEdited This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	// Optional. New name of the topic, if it was edited
	Name string `json:"name,omitempty"`
	// Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicReopened This object represents a service message about a forum topic reopened in the chat. Currently holds no information.
type ForumTopicReopened struct {
}

// GeneralForumTopicHidden This object represents a service message about General forum topic hidden in the chat. Currently holds no information.
type GeneralForumTopicHidden struct {
}

// GeneralForumTopicUnhidden This object represents a service message about General forum topic unhidden in the chat. Currently holds no information.
type GeneralForumTopicUnhidden struct {
}

type UserShared struct {
	// Identifier of the request
	RequestId int64 `json:"request_id"`
	// Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means.
	UserId int64 `json:"user_id"`
}

type ChatShared struct {
	// Identifier of the request
	RequestId int64 `json:"request_id"`
	// Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
	ChatId int64 `json:"chat_id"`
}

// WriteAccessAllowed This object represents a service message about a user allowing a bot to write messages after adding the bot to the attachment menu or launching a Web App from a link.
type WriteAccessAllowed struct {
	// Optional. Name of the Web App which was launched from a link
	WebAppName string `json:"web_app_name,omitempty"`
}
