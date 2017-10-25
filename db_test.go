package npidb_test

import (
	"database/sql"
	"errors"
	"time"

	. "github.com/cwarden/npidb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type queryResult int

const (
	noRows queryResult = iota
	queryError
	validDate
)

type stubDbUpdater struct {
	result queryResult
}

func (u stubDbUpdater) GetLastUpdate() (lastDate time.Time, err error) {
	switch u.result {
	case noRows:
		err = sql.ErrNoRows
	case queryError:
		err = errors.New("Query Failed")
	case validDate:
		lastDate = time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	return
}

func (u stubDbUpdater) RecordUpdate(t time.Time) error {
	return nil
}

var _ = Describe("Test", func() {
	var (
		stub stubDbUpdater
	)

	Describe("LastUpdate", func() {
		It("should return GetLastUpdate", func() {
			stub = stubDbUpdater{result: validDate}
			timeStamp, err := LastUpdate(stub)
			Expect(err).ToNot(HaveOccurred())
			Expect(timeStamp).To(Equal(time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)))
		})
		It("should return error if GetLastUpdate fails", func() {
			stub = stubDbUpdater{result: queryError}
			_, err := LastUpdate(stub)
			Expect(err).To(HaveOccurred())
		})
		It("should return default date if GetLastUpdate returns sql.ErrNoRows", func() {
			stub = stubDbUpdater{result: noRows}
			timeStamp, err := LastUpdate(stub)
			Expect(err).ToNot(HaveOccurred())
			Expect(timeStamp).To(Equal(time.Date(2017, time.June, 11, 0, 0, 0, 0, time.UTC)))
		})
	})
})
