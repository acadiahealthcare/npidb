package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cwarden/npidb"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: npi-update <sqlserver://username:password@server:port?database=dbname>")
		os.Exit(1)
	}
	connectionString := os.Args[1]
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	u := npidb.DbUpdateKeeper{db}
	lastDate, err := npidb.LastUpdate(u)
	if err != nil {
		log.Fatal(err)
	}
	updateFile, endDate := npidb.NextUrl(lastDate)
	fmt.Println("getting NPI updates: " + updateFile)

	csv, err := npidb.GetCSV(updateFile)
	if err != nil {
		log.Fatal(err)
	}
	err = npidb.WriteCsvToDb(csv, db, "NPI", "NPI")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Recording updates")
	err = npidb.RecordUpdate(u, endDate)
	if err != nil {
		log.Fatalf("Failed to record last update in NPI_Update table: %s", err)
	}
}
