package dbinterface

import (
	"fmt"
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter6/database"

)

func Query(db DB) error {
	name := "Aaron"
	rows, err := db.Query(`select name, created from example where name=$1`, name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e database.Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return nil
		}
		fmt.Printf("Results: \n\tName: %s\n\tCreated: %v\n ", e.Name, e.Created)
	}
	return rows.Err()
}

