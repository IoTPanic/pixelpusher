package db

import "fmt"

// DropAllTables will drop, all of the tables. Shocker.
func DropAllTables() error {
	for _, i := range []string{"users", "devices", "channels", "matrixes"} {
		err := DeleteTable(i)
		if err != nil {
			return err
		}
	}
	return nil
}

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
