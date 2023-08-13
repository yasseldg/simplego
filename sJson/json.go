package sJson

import (
	"encoding/json"
	"io"
	"os"

	"github.com/yasseldg/simplego/sLog"
)

func ToObj(msg string, obj any) error {
	err := json.Unmarshal([]byte(msg), obj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(msg), obj): %s", err)
		return err
	}
	return nil
}

func ToJson(v interface{}) (string, error) {
	result, err := json.Marshal(v)
	if err != nil {
		sLog.Error("json.Marshal(v): %s", err)
		return "", err
	}
	return string(result), nil
}

func Export(path string, objects any) error {
	f, err := os.Create(path)
	if err != nil {
		sLog.Error("ExportJson: os.OpenFile( %s ) : %s", path, err)
		return err
	}
	defer f.Close()

	data, err := json.MarshalIndent(objects, "", " ")
	if err != nil {
		sLog.Error("ExportJson: json.MarshalIndent(objects): %s : %s", path, err)
		return err
	}

	b, err := f.Write(data)
	if err != nil {
		sLog.Error("ExportJson: f.Write(data): %s : %s", path, err)
		return err
	}

	bb, err := f.WriteString("\n")
	if err != nil {
		sLog.Error("ExportJson: f.WriteString(): path: %s : %s", path, err)
		return err
	}

	sLog.Info("ExportJson: %d bytes written successfully in ( %s ). \n", (b + bb), path)
	return nil
}

func Import(path string, objects any) error {
	f, err := os.Open(path)
	if err != nil {
		sLog.Error("ImportJson: os.Open( %s ): %s", path, err)
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		sLog.Error("ImportJson: io.ReadAll(f): %s", err)
		return err
	}

	err = json.Unmarshal(data, objects)
	if err != nil {
		sLog.Error("ImportJson: json.Unmarshal(data, objects): %s", err)
		return err
	}

	sLog.Info("ImportJson: successful import of ( %s ). \n", path)
	return nil
}
