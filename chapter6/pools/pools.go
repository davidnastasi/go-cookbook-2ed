package pools

import (
	"database/sql"
	"fmt"
	"os"
)

func Setup() (*sql.DB, error) {
	pqinfo := fmt.Sprintf("host=localhost user=%s dbname=gocookbook sslmode=disable", os.Getenv("POSTGRESQL_USERNAME"))
	db, err := sql.Open("postgres", pqinfo)
	if err != nil {
		return nil, err
	}
	// there will only ever be 24 open connections
	db.SetMaxOpenConns(24)
	// MaxIdleConns can never be less than max open
	// SetMaxOpenConns otherwise it'll default to that value
	db.SetMaxIdleConns(24)
	return db, nil
}
