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
		return "", err
	}
	return string(result), nil
}

func Export(path string, objects any) (err error) {

	f, err := os.Create(path)

	if err == nil {
		defer f.Close()

		data, err := json.MarshalIndent(objects, "", " ")
		if err == nil {

			b, err := f.Write(data)
			if err == nil {

				bb, err := f.WriteString("\n")
				if err == nil {
					sLog.Info("ExportJson: %d bytes written successfully in ( %s ). \n", (b + bb), path)
				} else {
					sLog.Error("ExportJson: f.WriteString(): path: %s : %s", path, err)
				}
			} else {
				sLog.Error("ExportJson: f.Write(data): %s : %s", path, err)
			}
		} else {
			sLog.Error("ExportJson: json.MarshalIndent(objects): %s : %s", path, err)
		}
	} else {
		sLog.Error("ExportJson: os.OpenFile( %s ) : %s", path, err)
	}

	return
}

func Import(path string, objects any) (err error) {

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()

		data, err := io.ReadAll(f)
		if err == nil {

			err := json.Unmarshal(data, objects)
			if err == nil {
				sLog.Info("ImportJson: successful import of ( %s ). \n", path)
			} else {
				sLog.Error("ImportJson: json.Unmarshal(data, objects): %s", err)
			}
		} else {
			sLog.Error("ImportJson: io.ReadAll(f): %s", err)
		}
	} else {
		sLog.Error("ImportJson: os.Open( %s ): %s", path, err)
	}

	return
}
