package pusher

import (
	"log"

	"github.com/IoTPanic/pixelpusher/internal/db"
)

func HandleActivation(request ActivationRequest) (ActivationResponseSuccess, error) {
	dev, err := db.QueryFixture(request.LongID)
	if err != nil {
		log.Printf("Failed to activate device %s - %s", request.LongID, err.Error())
		return ActivationResponseSuccess{}, err
	}
	ch, err := db.QueryFixtureChannels(request.LongID)
	f := CastFixture(dev, ch)
	log.Printf("Activating device %s", f.Name)
	// TODO UPDATE CHANNELS IN DB
	resp := ActivationResponseSuccess{Success: true, Session: 0, PixelID: 1, DevNonce: request.DevNonce}
	return resp, nil
}
