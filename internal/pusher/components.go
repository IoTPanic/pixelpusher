package pusher

type FixtureSpan []Fixture

// Fixture is a device with LED strips attached
type Fixture struct {
	Name       string     `json:"name"`
	DBID       int        `json:"omitempty"`
	LongID     string     `json:"ID"`
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

//
type Universe struct {
	Name     string    `json:"name"`
	ID       string    `json:"id"`
	Fixtures []Fixture `json:"fixtures"`
}

// THIS SECTION IS FOR MESSAGE TYPES USED IN THE MQTT COMMUNICATION WITH DEVICES

// ActivationRequest is a request from a fixture to join and come online
type ActivationRequest struct {
	LongID           string     `json:"ID"`
	Connection       Connection `json:"connection"`
	Channels         []Channel  `json:"channels"`
	Model            string     `json:"model"`
	ConnectionMethod string     `json:"method"`
	DevNonce         int        `json:"nonce"`
}

// // ActivationResponseSuccess is a success message to an activation
// type ActivationResponseSuccess struct {
// 	Success  bool `json:"success"`
// 	Session  int  `json:"session"`
// 	PixelID  int  `json:"pixelID"`
// 	DevNonce int  `json:"nonce"`
// }

// // ActivationResponseFailure is the activation failure message
// type ActivationResponseFailure struct {
// 	Success bool   `json:"success"`
// 	Reason  string `json:"reason"`
// 	Hold    bool   `json:"hold"`
// }

// ActivationResponse is a JSON document sent in response to an activation
// request, if success is true, then the session, fixtureID, and nonce will
// be present, else, reason and hold will be present.
type ActivationResponse struct {
	Success   bool   `json:"success"`
	Session   int    `json:"session"`
	FixtureID int    `json:"fixture"`
	Nonce     int    `json:"nonce"`
	Reason    string `json:"reason"`
	Hold      bool   `json:"bool"`
}
