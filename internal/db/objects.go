package db

type Fixture struct {
	ID               int64  `sql:"id"`
	Name             string `sql:"name"`
	LongID           string `sql:"longID"`
	PixelsID         int    `sql:"pixelsID"`
	UniverseID       int    `sql:"universeID"`
	ConnectionMethod string `sql:"connectionMethod"`
	ConnectionHost   string `sql:"connectionHost"`
	ConnectionPort   int    `sql:"connectionPort"`
	ConnectionMQTT   bool   `sql:"connectionWMQTT"`
}

type Channel struct {
	ID       int  `sql:"id"`
	Channel  int  `sql:"channel"`
	DeviceID int  `sql:"deviceID"`
	RGBW     bool `sql:"RGBW"`
	Length   int  `sql:"length"`
}

type Universe struct {
	ID   int    `sql:"id"`
	Name string `sql:"name"`
}
