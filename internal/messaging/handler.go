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
	case ActivationReqest:
		r := pusher.ActivationRequest{}
		err = json.Unmarshal(m.Message, &r)
		log.Println("Fixture attempting activation")
		if err != nil {
			log.Println("[ ERROR ]", err.Error())
			return
		}
		resp := message{Type: ActivationResponse}
		ar, err := pusher.HandleActivation(r)
		if err != nil {
			afr := pusher.ActivationResponseFailure{Success: false, Hold: false, Reason: err.Error()}
			resp.Message, err = json.Marshal(afr)
			if err != nil {
				panic(err)
			}
		} else {
			resp.Message, err = json.Marshal(ar)
		}
		pyld, _ := json.Marshal(resp)
		// Send downstream
		DownstreamChannel <- channelMessage{Device: r.LongID, Payload: pyld}
		break
	}
}
