package database

import "database/sql"

func Exec(db *sql.DB) error {
	defer db.Exec("drop table example")

	if err := Create(db); err != nil {
		return err
	}
	if err := Query(db, "Aaron"); err != nil {
		return nil
	}
	return nil
}


