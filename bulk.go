package npidb

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/paulbellamy/ratecounter"
)

func WriteCsvToDb(reader io.Reader, txn *sql.Tx, tableName string) (headers []string, err error) {
	csvReader := csv.NewReader(reader)
	headers, err = csvReader.Read()
	if err != nil {
		return
	}

	bulk := mssql.CopyIn(tableName, mssql.MssqlBulkOptions{}, headers...)
	stmt, err := txn.Prepare(bulk)
	if err != nil {
		return
	}

	counter := ratecounter.NewRateCounter(60 * time.Second)

	fmt.Printf("Loading into %s\n", tableName)
	for {
		var row []string
		row, err = csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		fields := make([]interface{}, len(row))
		for i, v := range row {
			fields[i] = v
		}
		_, err = stmt.Exec(fields...)
		if err != nil {
			return
		}

		counter.Incr(1)
		fmt.Printf("\r%d rows/s", counter.Rate()/60)
	}

	fmt.Printf("\nFinishing write to %s\n", tableName)

	result, err := stmt.Exec()
	if err != nil {
		return
	}

	err = stmt.Close()
	if err != nil {
		return
	}

	rowCount, _ := result.RowsAffected()
	fmt.Printf("%d rows inserted\n", rowCount)
	return
}
