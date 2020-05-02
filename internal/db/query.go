package db

import (
	"errors"
	"fmt"
)

// QueryAllDevices returns every device that is registered in the database
func QueryAllDevices() ([]Device, error) {
	var result []Device
	rows, err := db.Query("SELECT * FROM devices")
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var d Device
		rows.Scan(&d.ID, &d.Name, &d.Project, &d.LongID, &d.Hostname, &d.Port, &d.Connector, &d.Key, &d.UseKey)
		result = append(result, d)
	}
	return result, nil
}

// QueryDevice by the SQL ID
func QueryDevice(deviceID int64) (Device, error) {
	var result Device
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM 'devices' WHERE ID=\"%d\"", deviceID))
	if err != nil {
		return result, err
	}
	iterated := false
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.Project, &result.LongID, &result.Hostname, &result.Port, &result.Connector, &result.Key, &result.UseKey)
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

func QueryUserFromUsername(username string) (User, error) {
	var result User
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM 'users' WHERE username=\"%s\"", username))
	if err != nil {
		return result, err
	}
	err = rows.Scan(result.ID, &result.Name, &result.Username, &result.Password, &result.Created, &result.LastLogin)
	return result, nil
}

func QueryUserFromID(id int64) (User, error) {
	var result User
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM 'users' WHERE id=\"%d\"", id))
	if err != nil {
		return result, err
	}
	err = rows.Scan(result.ID, &result.Name, &result.Username, &result.Password, &result.Created, &result.LastLogin)
	return result, nil
}
