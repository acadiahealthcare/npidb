package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cwarden/npidb"
	"github.com/cwarden/npidb/schema"
)

func main() {
	var dest string
	flag.StringVar(&dest, "table", "NPI", "name of table to update")
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Println("Usage: npi-init <sqlserver://username:password@server:port?database=dbname> /path/to/NPPES_Data_Dissemination_....zip")
		os.Exit(1)
	}
	connectionString := flag.Arg(0)
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	initFile := flag.Arg(1)
	fmt.Println("loading initial NPI file: " + initFile)

	reader, err := npidb.GetCSV(fmt.Sprintf("file://%s", initFile))
	if err != nil {
		log.Fatal(err)
	}

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	err = schema.CreateTables(txn, dest)
	if err != nil {
		log.Fatal(err)
	}
	_, err = npidb.WriteCsvToDb(reader, txn, dest)
	if err != nil {
		log.Fatal(err)
	}
	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
