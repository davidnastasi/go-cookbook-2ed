package dbinterface

func Create(db DB) error {
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATE)"); err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO example (name, created) values ('Aaron', NOW())`); err != nil {
		return err
	}
	return nil
}