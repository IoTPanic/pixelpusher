package pusher

import (
	"log"

	"github.com/IoTPanic/pixelpusher/internal/db"
)

func HandleActivation(request ActivationRequest, device string) error {
	dev, err := db.QueryFixture(request.LongID)
	if err != nil {
		log.Printf("Failed to activate device %s - %s", request.LongID, err.Error())
		return err
	}
	ch, err := db.QueryFixtureChannels(request.LongID)
	f := CastFixture(dev, ch)
	log.Printf("Activating device %s", f.Name)
	return nil
}
