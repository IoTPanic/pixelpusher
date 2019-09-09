package messaging

import (
	"encoding/json"
)

const (
	ActivationReqest   = 0
	ActivationResponse = 1
	DetailUpdate       = 2
	ControlPacket      = 3
	ApplicationData    = 4
)

type message struct {
	Type    int             `json:"type"`
	Message json.RawMessage `json:"message"`
}

type channelMessage struct {
	Payload []byte
	Device  string
}
