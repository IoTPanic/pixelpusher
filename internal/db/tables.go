package db

var createFixturesTableStatement = `CREATE TABLE IF NOT EXISTS fixtures (
	id INTEGER PRIMARY KEY,
	name VARCHAR(100),
	longID VARCHAR(100),
	pixelsID INT,
	connectionMethod VARCHAR(10),
	connectionHost VARCHAR(50),
	connectionPort INT,
	connectionWMQTT BOOLEAN
)`

var createChannelsTableStatement = `CREATE TABLE IF NOT EXISTS channels (
	id INTEGER PRIMARY KEY,
	channel INT,
	deviceID VARCHAR(100),
	RGBW BOOLEAN,
	length INT
)
`
