package main

import (
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter6/database"
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter6/dbinterface"
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// this wont do anything if commit is successful
	defer tx.Rollback()

	if err := dbinterface.Exec(tx); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

