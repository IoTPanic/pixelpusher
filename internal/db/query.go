package db

import "strconv"

func QueryFixtures() ([]Fixture, error) {
	var result []Fixture
	rows, err := db.Query("SELECT * FROM fixtures")
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var f Fixture
		rows.Scan(&f.ID, &f.LongID, &f.Name, &f.PixelsID, &f.ConnectionHost, &f.ConnectionPort, &f.ConnectionMethod, &f.ConnectionMQTT)
		result = append(result, f)
	}
	return result, nil
}

func QueryFixtureChannels(deviceID int) ([]Channel, error) {
	var result []Channel
	rows, err := db.Query("SELECT * FROM channels WHERE deviceID=" + strconv.Itoa(deviceID))
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var c Channel
		err = rows.Scan(&c.ID, &c.DeviceID, &c.Channel, &c.Length, &c.RGBW)
		result = append(result, c)
		if err != nil {
			return []Channel{}, err
		}
	}
	return result, err
}
