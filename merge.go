package npidb

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang-sql/sqlexp"
)

func MergeCsvToDb(reader io.Reader, db *sql.DB, tableName string, keyField string) (err error) {
	quoter, _ := sqlexp.QuoterFromDriver(db.Driver(), context.TODO())

	txn, err := db.Begin()
	if err != nil {
		return
	}

	tempTableName := "#Merge_Source"
	_, err = txn.Query(fmt.Sprintf("SELECT * INTO %s FROM %s WHERE 1 <> 1", quoter.ID(tempTableName), quoter.ID(tableName)))
	if err != nil {
		return
	}

	headers, err := WriteCsvToDb(reader, txn, tempTableName)
	if err != nil {
		return
	}

	merge := buildMergeSQL(headers, tableName, tempTableName, keyField, quoter)
	fmt.Printf("Merging into %s\n", tableName)
	_, err = txn.Query(merge)
	if err != nil {
		return
	}
	err = txn.Commit()
	return
}

func buildMergeSQL(cols []string, tableName string, tempTableName string, keyField string, quoter sqlexp.Quoter) (mergeSQL string) {
	quotedCols := quoteCols(cols, quoter)
	var updates, sourceCols []string
	for _, col := range quotedCols {
		updates = append(updates, fmt.Sprintf("target.%[1]s = source.%[1]s", col))
		sourceCols = append(sourceCols, fmt.Sprintf("source.%s", col))
	}
	mergeOn := fmt.Sprintf("target.%[1]s = source.%[1]s", quoter.ID(keyField))
	mergeSQL = fmt.Sprintf(`
		MERGE
			%s AS target
		USING %s AS source
		ON
			%s
		WHEN MATCHED THEN UPDATE SET
			%s
		WHEN NOT MATCHED THEN
			INSERT(%s)
			VALUES(%s);`,
		quoter.ID(tableName), quoter.ID(tempTableName), mergeOn, strings.Join(updates, ","), strings.Join(quotedCols, ","), strings.Join(sourceCols, ","))
	return
}

func quoteCols(columns []string, quoter sqlexp.Quoter) []string {
	var quotedCols []string
	for _, col := range columns {
		quotedCols = append(quotedCols, quoter.ID(col))
	}
	return quotedCols
}
