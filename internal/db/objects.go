package db

import "time"

// Device pixelpusher object for DB
type Device struct {
	ID        int64  `sql:"id"`
	Name      string `sql:"name"`
	LongID    string `sql:"longID"`
	Hostname  string `sql:"hostname"`
	Port      int    `sql:"port"`
	Connector bool   `sql:"connector"`
	Key       string `sql:"key"`
	UseKey    bool   `sql:"use_key"`
}

// Channel pixelpusher object for DB
type Channel struct {
	ID        int    `sql:"id"`
	Name      string `sql:"name"`
	Type      int    `sql:"type"`
	Device    int    `sql:"device"`
	Color     string `sql:"color"`
	MaxLength int    `sql:"max_length"`
}

// Matrix pixelpusher object for DB
type Matrix struct {
	ID         int    `sql:"id"`
	Device     int    `sql:"device"`
	Channel    int    `sql:"channel"`
	Width      int    `sql:"width"`
	Height     int    `sql:"height"`
	Type       int    `sql:"type"`
	Coloring   string `sql:"coloring"`
	Offset     int    `sql:"offset"`
	Brightness int    `sql:"brightness"`
}

// Project pixelpusher object for DB
type Project struct {
	ID            int       `sql:"id"`
	Name          string    `sql:"name"`
	Created       time.Time `sql:"created"`
	LastUpdate    time.Time `sql:"last_update"`
	Client        string    `sql:"client"`
	Active        bool      `sql:"active"`
	FrontendState string    `sql:"frontend_state"`
}

// User pixelpusher object for DB
type User struct {
	ID        int       `sql:"id"`
	Name      string    `sql:"name"`
	Username  string    `sql:"username"`
	Password  string    `sql:"password"`
	Created   time.Time `sql:"created"`
	LastLogin time.Time `sql:"last_login"`
}
