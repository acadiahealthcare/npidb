package npidb

import (
	"database/sql"
	"fmt"
	"time"
)

type UpdateKeeper interface {
	GetLastUpdate() (time.Time, error)
	RecordUpdate(time.Time) error
}

type DbUpdateKeeper struct {
	DB        *sql.DB
	TableName string
}

func (u DbUpdateKeeper) GetLastUpdate() (lastDate time.Time, err error) {
	err = u.DB.QueryRow(fmt.Sprintf("SELECT TOP 1 end_date FROM %s ORDER BY end_date DESC", u.TableName)).Scan(&lastDate)
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
	_, err = u.DB.Exec(fmt.Sprintf("INSERT INTO %s (end_date) VALUES (?)", u.TableName), t)
	return
}

func RecordUpdate(u UpdateKeeper, t time.Time) error {
	return u.RecordUpdate(t)
}
