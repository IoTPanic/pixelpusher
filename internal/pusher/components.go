package pusher

// Fixture is a end device that controls LED pixels
type Fixture struct {
	Name     string
	Channels []Channel
	ID       uint16
}

// Channel is a object to represent a single line of pixels on a fixture, a fixture can have multiple channels
type Channel struct {
	RGBW   bool
	Pixels int
}

// Universe is a collection of devices that will receive the same data
type Universe struct {
	Name     string
	Fixtures []Fixture
}
