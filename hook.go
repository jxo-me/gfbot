package telebot

import (
	"encoding/json"
	"net/http"
)

// IHook is a provider of Webhook.
type IHook interface {
	GetFiles() map[string]File
	GetParams() map[string]string
	WaitForStop(stop chan struct{})
	Handler(w http.ResponseWriter, r *http.Request)
}

// Webhook returns the current webhook status.
func (b *Bot) Webhook() (*IHook, error) {
	data, err := b.Raw("getWebhookInfo", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result IHook
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return &resp.Result, nil
}

// SetWebhook configures a bot to receive incoming
// updates via an outgoing webhook.
func (b *Bot) SetWebhook(w IHook) error {
	_, err := b.sendFiles("setWebhook", w.GetFiles(), w.GetParams())
	return err
}

// RemoveWebhook removes webhook integration.
func (b *Bot) RemoveWebhook(dropPending ...bool) error {
	drop := false
	if len(dropPending) > 0 {
		drop = dropPending[0]
	}
	_, err := b.Raw("deleteWebhook", map[string]bool{
		"drop_pending_updates": drop,
	})
	return err
}
