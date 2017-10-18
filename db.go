package npidb

import (
	"database/sql"
	"fmt"
	"time"
)

// Get the date of the last weekly update saved to the database
func LastUpdate(db *sql.DB) (lastDate time.Time, err error) {
	err = db.QueryRow("SELECT TOP 1 end_date FROM NPI_Update ORDER BY end_date DESC").Scan(&lastDate)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("using default date")
		lastDate = time.Date(2017, time.June, 11, 0, 0, 0, 0, time.UTC)
	case err != nil:
		return
	default:
	}
	return
}

func RecordUpdate(db *sql.DB, t time.Time) (sql.Result, error) {
	return db.Exec("INSERT INTO NPI_Update (end_date) VALUES (?)", t)
}
