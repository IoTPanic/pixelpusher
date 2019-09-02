package db

type Fixture struct {
	ID               int    `sql:"id"`
	Name             string `sql:"name"`
	LongID           string `sql:"longID"`
	PixelsID         int    `sql:"pixelsID"`
	ConnectionMethod string `sql:"connectionMethod"`
	ConnectionHost   string `sql:"connectionHost"`
	ConnectionPort   int    `sql:"connectionPort"`
	ConnectionMQTT   bool   `sql:"connectionWMQTT"`
}

type Channel struct {
	ID       int    `sql:"id"`
	Channel  int    `sql:"channel"`
	DeviceID string `sql:"deviceID"`
	RGBW     bool   `sql:"RGBW"`
	Length   int    `sql:"length"`
}
