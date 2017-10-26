package npidb

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func NextUrl(after time.Time) (updateFile string, endDate time.Time) {
	startDate := after.AddDate(0, 0, 1)
	endDate = after.AddDate(0, 0, 7)
	dateLayout := "010206"
	updateFile = fmt.Sprintf("http://download.cms.gov/nppes/NPPES_Data_Dissemination_%s_%s_Weekly.zip", startDate.Format(dateLayout), endDate.Format(dateLayout))
	return
}

// Fetch an NPI weekly update zip file and return an io.Reader for the csv file
// within.  File can be fetched from remote host or using file:// URI.
func GetCSV(url string) (csvReader io.Reader, err error) {
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c := &http.Client{Transport: t}
	resp, err := c.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode >= 400 {
		err = errors.New(fmt.Sprintf("Could not retrieve %s: %s", url, resp.Status))
		return
	}
	defer resp.Body.Close()

	tmpfile, err := ioutil.TempFile("", "npi-update")
	if err != nil {
		return
	}
	defer os.Remove(tmpfile.Name())
	length, err := io.Copy(tmpfile, resp.Body)
	if err != nil {
		return
	}

	r, err := zip.NewReader(tmpfile, length)
	if err != nil {
		return
	}

	for _, f := range r.File {
		if !strings.HasSuffix(f.Name, ".csv") || strings.HasSuffix(f.Name, "Header.csv") {
			continue
		}
		csvReader, err = f.Open()
		if err != nil {
			return
		}
		return
	}
	err = errors.New("NPI update not found in " + url)
	return
}
