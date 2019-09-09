package messaging

import (
	"encoding/json"
)

const (
	activationRequestMessage = 0x00
)

type message struct {
	Type    int             `json:"type"`
	Message json.RawMessage `json:"message"`
}
