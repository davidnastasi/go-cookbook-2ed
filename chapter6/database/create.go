package database

import "database/sql"

func Create(db *sql.DB) error {
	if _, err := db.Exec("create table example (name varchar(20), created date)"); err != nil {
		return err
	}
	if _, err := db.Exec(`insert into example (name, created) values ('Aaron', now())`); err != nil {
		return err
	}

	return nil
}
