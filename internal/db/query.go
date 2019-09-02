package db

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
