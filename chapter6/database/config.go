package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type Example struct {
	Name string
	Created *time.Time
}

func Setup() (*sql.DB, error) {
	//os.Getenv("POSTGRESQL_USERNAME")
	//os.Getenv("POSTGRESQL_PASSWORD")
	pqinfo := fmt.Sprintf("host=localhost user=%s dbname=gocookbook sslmode=disable", os.Getenv("POSTGRESQL_USERNAME"))
	db, err := sql.Open("postgres",pqinfo)
	if err != nil {
		return nil, err
	}
	return db, nil

}

