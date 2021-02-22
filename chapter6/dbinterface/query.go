package database

import (
	"database/sql"
	"fmt"
)

func Query(db *sql.DB, name string) error {
	rows, err := db.Query(`select  name, created from example where name=$1 `, name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return nil
		}
		fmt.Printf("Results: \n\tName: %s\n\t Created: %v\n ", e.Name, e.Created)
	}
	return rows.Err()
}

