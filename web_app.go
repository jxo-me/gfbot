package telebot

import "time"

// WebApp represents a parameter of the inline keyboard button
// or the keyboard button used to launch Web App.
type WebApp struct {
	// An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
	URL string `json:"url"`
}

// WebAppMessage describes an inline message sent by a Web App on behalf of a user.
type WebAppMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}

// WebAppData object represents a data sent from a Web App to the bot
type WebAppData struct {
	Data string `json:"data"`
	Text string `json:"button_text"`
}

// WebAppChat describes chat information:
// https://core.telegram.org/bots/webapps#webappchat
type WebAppChat struct {
	// Unique identifier for this chat.
	ID int64 `json:"id"`

	// Type of chat.
	Type ChatType `json:"type"`

	// Title of the chat.
	Title string `json:"title"`

	// Optional. Username of the chat.
	Username string `json:"username,omitempty"`

	// Optional. URL of the chat’s photo. The photo can be in .jpeg or .svg
	// formats. Only returned for Web Apps launched from the attachment menu.
	PhotoURL string `json:"photo_url,omitempty"`
}

// WebAppUser describes user information:
// https://core.telegram.org/bots/webapps#webappuser
type WebAppUser struct {
	// A unique identifier for the user or bot.
	ID int64 `json:"id"`

	// Optional. True, if this user is a bot. Returned in the `receiver` field
	// only.
	IsBot bool `json:"is_bot,omitempty"`

	// First name of the user or bot.
	FirstName string `json:"first_name"`

	// Optional. Last name of the user or bot.
	LastName string `json:"last_name,omitempty"`

	// Optional. Username of the user or bot.
	Username string `json:"username,omitempty"`

	// Optional. IETF language tag of the user's language. Returns in user
	// field only.
	//
	// See: https://en.wikipedia.org/wiki/IETF_language_tag
	LanguageCode string `json:"language_code,omitempty"`

	// Optional. True, if this user is a Telegram Premium user.
	IsPremium bool `json:"is_premium,omitempty"`

	// Optional. URL of the user’s profile photo. The photo can be in .jpeg or
	// .svg formats. Only returned for Web Apps launched from the
	// attachment menu.
	PhotoURL string `json:"photo_url,omitempty"`
}

// WebAppInitData describes parsed initial data sent from TWA application. You can
// find specification for all the parameters in the official documentation:
// https://core.telegram.org/bots/webapps#webappinitdata
type WebAppInitData struct {
	// A unique identifier for the Web App session, required for sending
	// messages via the answerWebAppQuery method.
	//
	// See: https://core.telegram.org/bots/api#answerwebappquery
	QueryID string `json:"query_id,omitempty"`

	// An object containing data about the current user.
	User *WebAppUser `json:"user,omitempty"`

	// An object containing data about the chat partner of the current user in
	// the chat where the bot was launched via the attachment menu.
	// Returned only for private chats and only for Web Apps launched
	// via the attachment menu.
	Receiver *WebAppUser `json:"receiver,omitempty"`

	// An object containing data about the chat where the bot was
	// launched via the attachment menu. Returned for supergroups, channels
	// and group chats – only for Web Apps launched via the attachment menu.
	Chat *WebAppChat `json:"chat,omitempty"`

	// Optional. Type of the chat from which the Web App was opened.
	// Can be either “sender” for a private chat with the user opening the link,
	// “private”, “group”, “supergroup”, or “channel”.
	// Returned only for Web Apps launched from direct links.
	ChatType ChatType `json:"chat_type,omitempty"`

	// Optional. Global identifier, uniquely corresponding to the chat from which the Web App was opened.
	// Returned only for Web Apps launched from a direct link.
	ChatInstance string `json:"chat_instance,omitempty"`

	// Optional. The value of the `startattach` parameter, passed via link. Only
	// returned for Web Apps when launched from the attachment menu via link.
	StartParam string `json:"start_param,omitempty"`

	// Optional. Time in seconds, after which a message can be sent via the
	// `answerWebAppQuery` method.
	//
	// See: https://core.telegram.org/bots/api#answerwebappquery
	CanSendAfterRaw int `json:"can_send_after,omitempty"`

	// Init data generation date.
	AuthDateRaw int `json:"auth_date"`
	// A hash of all passed parameters, which the bot server can use to
	// check their validity.
	//
	// See: https://core.telegram.org/bots/webapps#validating-data-received-via-the-web-app
	Hash string `json:"hash"`
}

// AuthDate returns AuthDateRaw as time.Time.
func (d *WebAppInitData) AuthDate() time.Time {
	return time.Unix(int64(d.AuthDateRaw), 0)
}

// CanSendAfter returns computed time which depends on CanSendAfterRaw and
// AuthDate. Originally, CanSendAfterRaw means time in seconds, after which
// `answerWebAppQuery` method can be called and that's why this value could
// be computed as time.
func (d *WebAppInitData) CanSendAfter() time.Time {
	return d.AuthDate().Add(time.Duration(d.CanSendAfterRaw) * time.Second)
}
