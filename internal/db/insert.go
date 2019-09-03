package db

func (f Fixture) Insert() (int64, error) {
	stmt, err := db.Prepare("INSERT INTO fixtures(name, longID, pixelsID, connectionMethod, connectionHost, connectionPort, connectionWMQTT, universeID) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(f.Name, f.LongID, f.PixelsID, f.ConnectionMethod, f.ConnectionHost, f.ConnectionPort, f.ConnectionMQTT, f.UniverseID)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (c Channel) Insert() (int64, error) {
	stmt, err := db.Prepare("INSERT INTO channels(channel, deviceID, RGBW, length)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(c.Insert, c.DeviceID, c.RGBW, c.Length)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}
