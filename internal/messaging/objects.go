package messaging

import (
	"encoding/json"

	"github.com/IoTPanic/pixelpusher/internal/pusher"
)

type message struct {
	Type    int             `json:"type"`
	Message json.RawMessage `json:"message"`
}

type activationRequest struct {
	Name       string            `json:"fixture-name"`
	LongID     string            `json:"longsID"`
	Connection pusher.Connection `json:"connection"`
	Channels   []pusher.Channel  `json:"channels"`
	Model      string            `json:"model"`
}
