package db

import "errors"

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

func QueryFixture(deviceID string) (Fixture, error) {
	var result Fixture
	rows, err := db.Query("SELECT id, name, longID, pixelsID, universeID, connectionMethod, connectionHost, connectionPort, connectionWMQTT FROM `fixtures` WHERE longID=\"" + deviceID + "\"")
	if err != nil {
		return result, err
	}
	iterated := false
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.LongID, &result.PixelsID, &result.UniverseID, &result.ConnectionMethod, &result.ConnectionHost, &result.ConnectionPort, &result.ConnectionMQTT)
		if iterated {
			return result, errors.New("duplicate long ID entries")
		}
		if err != nil {
			return result, err
		}
		iterated = true
	}
	return result, nil
}

func QueryFixtureChannels(deviceLongID string) ([]Channel, error) {
	var result []Channel
	rows, err := db.Query("SELECT * FROM channels WHERE deviceID=" + deviceLongID)
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

func QueryUniverses() ([]Universe, error) {
	var result []Universe
	rows, err := db.Query("SELECT * FROM universes")
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var u Universe
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}
