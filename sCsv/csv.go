package sCsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/yasseldg/mgm/v4"
	"github.com/yasseldg/simplego/sFile"
	"github.com/yasseldg/simplego/sLog"
	"github.com/yasseldg/simplego/sMongo"
)

func Import(file_path string) ([][]string, error) {
	var importFunc func(string) ([][]string, error)

	switch {
	case strings.HasSuffix(file_path, ".csv"):
		importFunc = importCsv
	case strings.HasSuffix(file_path, ".csv.gz"):
		importFunc = importGz
	default:
		return nil, errors.New("unsupported file format")
	}

	return importFunc(file_path)
}

func Export(file_path string, data [][]string) error {
	f, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		sLog.Error("Export: os.OpenFile( %s ): %s", file_path, err)
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		sLog.Error("Export: csvReader.WriteAll(data) in ( %s ): %s", file_path, err)
		return err
	}

	sLog.Info("Export: %d lines written successfully in ( %s )", len(data), file_path)
	return nil
}

func importCsv(file_path string) ([][]string, error) {
	file, err := os.Open(file_path)
	if err != nil {
		sLog.Error("Import: os.Open( %s ): %s", file_path, err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		sLog.Error("Import: reader.ReadAll(): %s", err)
		return nil, err
	}

	sLog.Info("Import: successful import of ( %s )", file_path)
	return data, nil
}

func importGz(file_path string) ([][]string, error) {
	new_file_path := sFile.DecompressGzip(file_path)
	if len(new_file_path) == 0 {
		err := fmt.Errorf("importGz: Failed to decompress ( %s )", file_path)
		sLog.Error(err.Error())
		return nil, err
	}
	defer sFile.DeletePath(new_file_path)

	return Import(new_file_path)
}

type SaveFunc func(objs []mgm.Model, coll sMongo.Collection) (err error)
type ObjFunc func(data []string) (obj mgm.Model)
type ImportFunc func(file_path string, obj_func ObjFunc, save_func SaveFunc, batch_size int, coll sMongo.Collection) (err error)

func ImportAndSave(file_path string, obj_func ObjFunc, save_func SaveFunc, batch_size int, coll sMongo.Collection) (err error) {
	var importFunc ImportFunc

	switch {
	case strings.HasSuffix(file_path, ".csv"):
		importFunc = importSaveCsv
	case strings.HasSuffix(file_path, ".csv.gz"):
		importFunc = importSaveGz
	default:
		return errors.New("unsupported file format")
	}

	return importFunc(file_path, obj_func, save_func, batch_size, coll)
}

func readBatch(reader *csv.Reader, batch_size int) ([][]string, error) {
	var lines [][]string
	for i := 0; i < batch_size; i++ {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func processBatch(lines [][]string, obj_func ObjFunc) []mgm.Model {
	var objs []mgm.Model
	for _, line := range lines {
		obj := obj_func(line)
		if obj == nil {
			continue
		}
		objs = append(objs, obj)
	}
	return objs
}

func importSaveCsv(file_path string, obj_func ObjFunc, save_func SaveFunc, batch_size int, coll sMongo.Collection) error {
	f, err := os.Open(file_path)
	if err != nil {
		sLog.Error("Import: os.Open( %s ): %s", file_path, err)
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	for {
		batch, err := readBatch(reader, batch_size)
		if err != nil {
			sLog.Error("Import: readBatch(reader): %s", err)
			return err
		}

		if len(batch) == 0 {
			break
		}

		objs := processBatch(batch, obj_func)
		if objs == nil {
			return fmt.Errorf("Import: processBatch(batch, obj_func)")
		}

		err = save_func(objs, coll)
		if err != nil {
			sLog.Error("Import: save_func(objs): %s", err)
			return err
		}
	}

	sLog.Info("Import: successful import of ( %s )", file_path)
	return nil
}

func importSaveGz(file_path string, obj_func ObjFunc, save_func SaveFunc, batch_size int, coll sMongo.Collection) error {
	new_file_path := sFile.DecompressGzip(file_path)
	if len(new_file_path) == 0 {
		err := fmt.Errorf("importGz: Failed to decompress ( %s )", file_path)
		sLog.Error(err.Error())
		return err
	}
	defer sFile.DeletePath(new_file_path)

	return ImportAndSave(new_file_path, obj_func, save_func, batch_size, coll)
}
