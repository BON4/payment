package dbtest

import (
	"database/sql"
	"fmt"
	"testing"
)

func ConnectTestDB(driver, constr string) (*sql.DB, error) {
	db, err := sql.Open(driver, constr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func FlushDB(t *testing.T, db *sql.DB, tabels ...string) error {
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	for _, table := range tabels {
		_, err = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}
