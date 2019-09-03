package pusher

type FixtureSpan []Fixture

// Fixture is a device with LED strips attached
type Fixture struct {
	Name       string     `json:"name"`
	DBID       int        `json:"omitempty"`
	LongID     string     `json:"longID"`
	PixelsID   uint16     `json:"pixelsID"`
	Connection Connection `json:"connection"`
	Channels   []Channel  `json:"channels"`
	UniverseID int        `json:"universe"`
}

// Connection is a UDP, TCP, or HTTP connection to a valid fixture
type Connection struct {
	IP   string `json:"IP"`
	Port int    `json:"port"`
	Type string `json:"type"`
	MQTT bool   `json:"mqtt"`
}

// Channel is a single strip
type Channel struct {
	ID         uint16 `json:"ID"`
	PixelCount uint16 `json:"pixelCnt"`
	RGBW       bool   `json:"RGBW"`
}
