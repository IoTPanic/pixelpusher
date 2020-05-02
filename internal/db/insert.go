package db

// Insert a device into the database
func (d Device) Insert() (int64, error) {
	stmt, err := db.Prepare("INSERT INTO fixtures(name, project, longID, hostname, port, connector, key, userkey) values(?,?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(d.Name, d.Project, d.LongID, d.Hostname, d.Port, d.Connector, d.Key, d.UseKey)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

// Insert a channel into the database
func (c Channel) Insert() (int64, error) {
	stmt, err := db.Prepare("INSERT INTO channels(name, project, type, device, color, max_length) values(?,?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(c.Name, c.Project, c.Type, c.Device, c.Color, c.MaxLength)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

// Insert a matrix into the database
func (m Matrix) Insert() (int64, error) {
	stmt, err := db.Prepare("INSERT INTO matrixes(device, channel, project, width, height, type, coloring, offset, brightness) values(?,?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(m.Device, m.Channel, m.Project, m.Width, m.Height, m.Type, m.Coloring, m.Offset, m.Brightness)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

// Insert project into the database
func (p Project) Insert() (int64, error) {
	stmt, err := db.Prepare("INSERT INTO projects(name, created, last_update, client, active, frontend_state) values(?,?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(p.Name, p.Created, p.LastUpdate, p.Client, p.Active, p.FrontendState)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

// Insert user into the database
func (u User) Insert() (int64, error) {
	stmt, err := db.Prepare("INSERT INTO projects(name, username, password, created, last_login) values(?,?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(u.Name, u.Username, u.Password, u.Created, u.LastLogin)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
