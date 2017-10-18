package npidb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dailyburn/ratchet"
	"github.com/dailyburn/ratchet/logger"
	"io"

	batch "github.com/cwarden/ratchet_batcher"
	merge "github.com/cwarden/ratchet_mergewriter"
	_ "github.com/denisenkom/go-mssqldb"
	procs "github.com/samuelhug/ratchet_processors"
)

func WriteCsvToDb(csv io.Reader, db *sql.DB, tableName string, keyField string) (err error) {
	logger.LogLevel = logger.LevelInfo
	var processorChain []ratchet.DataProcessor

	input, err := procs.NewCSVReader(csv)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error initializing input: %s", err))
		return
	}

	processorChain = append(processorChain, input)

	// We're constrained by the number of parameters, 2100, that can be
	// passed to SQL Server.
	batcher := batch.NewBatcher(6)
	processorChain = append(processorChain, batcher)

	output := merge.NewSQLMergeWriter(db, tableName, keyField)
	processorChain = append(processorChain, output)
	pipeline := ratchet.NewPipeline(processorChain...)
	err = <-pipeline.Run()
	if err != nil {
		err = errors.New("An error occurred in the data pipeline: " + err.Error())
	}
	return
}
