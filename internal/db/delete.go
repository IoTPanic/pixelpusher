package db

import "fmt"

// DeleteTable drops table from instance
func DeleteTable(tablename string) error {
	stmt, err := db.Prepare(fmt.Sprintln("DROP TABLE %s", tablename))
	if err != nil {
		return err
	}
	_, err = stmt.Exec()

	if err != nil {
		return err
	}
	return err
}
