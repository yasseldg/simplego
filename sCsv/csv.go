package sCsv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/yasseldg/simplego/sFile"
	"github.com/yasseldg/simplego/sLog"
)

func Import(path string) (data [][]string, err error) {

	f := func(string) (data [][]string, err error) { return }
	l := len(path)
	switch {
	case (path[l-4:] == ".csv"):
		f = file

	case (path[l-7:] == ".csv.gz"):
		f = fileGz
	}

	return f(path)
}

func Export(path string, data [][]string) (err error) {

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err == nil {
		defer f.Close()

		// write csv values using csv.Writer
		csvWriter := csv.NewWriter(f)
		defer csvWriter.Flush()

		err = csvWriter.WriteAll(data)
		if err == nil {
			sLog.Info("Export: %d lines written successfully in ( %s )", len(data), path)
		} else {
			sLog.Error("Export: csvReader.WriteAll(data) in ( %s ): %s", path, err)
		}
	} else {
		sLog.Error("Export: os.OpenFile( %s ): %s", path, err)
	}

	return
}

func file(path string) (data [][]string, err error) {

	// open f
	f, err := os.Open(path)
	if err == nil {
		// remember to close the f at the end of the program
		defer f.Close()

		// read csv values using csv.Reader
		csvReader := csv.NewReader(f)
		data, err = csvReader.ReadAll()
		if err == nil {
			sLog.Info("Import: successful import of ( %s )", path)
		} else {
			sLog.Error("Import: csvReader.ReadAll(): %s", err)
		}
	} else {
		sLog.Error("Import: os.Open( %s ): %s", path, err)
	}

	return
}

func fileGz(path string) (data [][]string, err error) {

	new_path := sFile.DecompressGzip(path)
	if len(new_path) > 0 {
		data, err = Import(new_path)
		if err != nil {
			sLog.Error("fileGz: Import( %s ): %s", new_path, err)
		}

		err = sFile.DeletePath(new_path)
	} else {
		err = fmt.Errorf("fileGz: decompressGzip( %s )", path)
	}
	return
}
