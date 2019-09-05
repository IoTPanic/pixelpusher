package pusher

import (
	"github.com/IoTPanic/pixelpusher/internal/db"
)

func CastFixture(f db.Fixture, c []db.Channel) Fixture {
	return Fixture{
		Name:     f.Name,
		LongID:   f.LongID,
		PixelsID: uint16(f.PixelsID),
		Channels: CastChannels(c),
		Connection: Connection{
			IP:   f.ConnectionHost,
			Port: f.ConnectionPort,
			Type: f.ConnectionMethod,
			MQTT: f.ConnectionMQTT,
		},
	}
}

func CastChannels(c []db.Channel) []Channel {
	result := make([]Channel, len(c))
	for i, o := range c {
		result[i] = Channel{PixelCount: uint16(o.Length), RGBW: o.RGBW, ID: uint16(o.Channel)}
	}
	return result
}

func (f Fixture) CastToDB() (db.Fixture, []db.Channel) {
	var channels []db.Channel
	result := db.Fixture{Name: f.Name, LongID: f.LongID, PixelsID: int(f.PixelsID), UniverseID: f.UniverseID, ConnectionHost: f.Connection.IP, ConnectionPort: f.Connection.Port, ConnectionMethod: f.Connection.Type, ConnectionMQTT: f.Connection.MQTT}
	for _, i := range f.Channels {
		channels = append(channels, i.CastToDB(f.DBID))
	}
	return result, channels
}

func (c Channel) CastToDB(deviceID int) db.Channel {
	return db.Channel{Channel: int(c.ID), DeviceID: deviceID, RGBW: c.RGBW, Length: int(c.PixelCount)}
}

func (u Universe) CastToDB() (db.Universe, []db.Fixture) {
	var r []db.Fixture
	for _, i := range u.Fixtures {
		r = append(r, db.Fixture{Name: i.Name, LongID: i.LongID, PixelsID: int(i.PixelsID), UniverseID: i.UniverseID, ConnectionMethod: i.Connection.Type, ConnectionHost: i.Connection.IP, ConnectionPort: i.Connection.Port, ConnectionMQTT: i.Connection.MQTT})
	}
	return db.Universe{Name: u.Name}, r
}
