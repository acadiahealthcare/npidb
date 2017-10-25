package npidb

import (
	"database/sql"
	"time"
)

type UpdateKeeper interface {
	GetLastUpdate() (time.Time, error)
	RecordUpdate(time.Time) error
}

type DbUpdateKeeper struct {
	DB *sql.DB
}

func (u DbUpdateKeeper) GetLastUpdate() (lastDate time.Time, err error) {
	err = u.DB.QueryRow("SELECT TOP 1 end_date FROM NPI_Update ORDER BY end_date DESC").Scan(&lastDate)
	return
}

// Get the date of the last weekly update saved to the database
func LastUpdate(u UpdateKeeper) (lastDate time.Time, err error) {
	lastDate, err = u.GetLastUpdate()
	switch {
	case err == sql.ErrNoRows:
		lastDate = time.Date(2017, time.June, 11, 0, 0, 0, 0, time.UTC)
		err = nil
	case err != nil:
		return
	default:
	}
	return
}

func (u DbUpdateKeeper) RecordUpdate(t time.Time) (err error) {
	_, err = u.DB.Exec("INSERT INTO NPI_Update (end_date) VALUES (?)", t)
	return
}

func RecordUpdate(u UpdateKeeper, t time.Time) error {
	return u.RecordUpdate(t)
}
