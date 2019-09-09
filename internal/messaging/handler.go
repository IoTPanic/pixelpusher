package messaging

import (
	"encoding/json"
	"log"

	"github.com/IoTPanic/pixelpusher/internal/pusher"
)

func consumeMessage(pyld []byte, topic string) {
	var m message
	// First, we're going to unmarshal the message, but really only the type will be unmarshalled
	err := json.Unmarshal(pyld, &m)
	if err != nil {
		log.Println("[ ERROR ] Failed to unmarshal message", err.Error())
	}
	switch m.Type {
	case activationRequestMessage:
		r := pusher.ActivationRequest{}
		err = json.Unmarshal(m.Message, &r)
		log.Println("Fixture attempting activation")
		if err != nil {
			log.Println("[ ERROR ]", err.Error())
			return
		}
		break
	}
}
