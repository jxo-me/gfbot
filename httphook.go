package telebot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type HttpHookEndpoint struct {
	PublicURL string `json:"public_url"`
}

type HttpHook struct {
	Listen         string   `json:"url"`
	MaxConnections int      `json:"max_connections"`
	AllowedUpdates []string `json:"allowed_updates"`
	IP             string   `json:"ip_address"`
	DropUpdates    bool     `json:"drop_pending_updates"`
	SecretToken    string   `json:"secret_token"`

	// (HttpHookInfo)
	HasCustomCert     bool   `json:"has_custom_certificate"`
	PendingUpdates    int    `json:"pending_update_count"`
	ErrorUnixtime     int64  `json:"last_error_date"`
	ErrorMessage      string `json:"last_error_message"`
	SyncErrorUnixtime int64  `json:"last_synchronization_error_date"`

	Endpoint *HttpHookEndpoint

	dest chan<- Update
	bot  *Bot
}

func (h *HttpHook) GetFiles() map[string]File {
	m := make(map[string]File)

	// check if it is overwritten by an endpoint
	if h.Endpoint != nil {
	}
	return m
}

func (h *HttpHook) GetParams() map[string]string {
	params := make(map[string]string)

	if h.MaxConnections != 0 {
		params["max_connections"] = strconv.Itoa(h.MaxConnections)
	}
	if len(h.AllowedUpdates) > 0 {
		data, _ := json.Marshal(h.AllowedUpdates)
		params["allowed_updates"] = string(data)
	}
	if h.IP != "" {
		params["ip_address"] = h.IP
	}
	if h.DropUpdates {
		params["drop_pending_updates"] = strconv.FormatBool(h.DropUpdates)
	}
	if h.SecretToken != "" {
		params["secret_token"] = h.SecretToken
	}

	if h.Endpoint != nil {
		params["url"] = h.Endpoint.PublicURL
	}
	return params
}

func (h *HttpHook) Poll(b *Bot, dest chan Update, stop chan struct{}) {
	if err := b.SetWebhook(h); err != nil {
		b.OnError(err, nil)
		close(stop)
		return
	}

	// store the variables so the HTTP-handler can use 'em
	h.dest = dest
	h.bot = b
}

func (h *HttpHook) WaitForStop(stop chan struct{}) {
	<-stop
	close(stop)
}

func (h *HttpHook) Handler(w http.ResponseWriter, r *http.Request) {
	if h.SecretToken != "" && r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != h.SecretToken {
		h.bot.debug(fmt.Errorf("invalid secret token in request"))
		return
	}

	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		h.bot.debug(fmt.Errorf("cannot decode update: %v", err))
		return
	}
	h.dest <- update
}

func (h *HttpHook) Signature(w http.ResponseWriter, r *http.Request) {
	initData := r.Header.Get("X-Telegram-Bot-Web-App-Authorization")
	ok, err := h.bot.ValidateWebAppData(initData, 0)
	if err != nil {
		h.bot.debug(fmt.Errorf("initData ValidateWebAppData error:%s", err.Error()))
		return
	}
	if ok {
		// process info into context Todo
	}

}
