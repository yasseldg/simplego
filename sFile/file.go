package sFile

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/yasseldg/simplego/sLog"
)

func FilesOnDir(dir_path string) []fs.FileInfo {

	entries, err := os.ReadDir(dir_path)
	if err != nil {
		sLog.Error("FilesOnDir os.ReadDir( %s ): %s", dir_path, err)
	}

	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			sLog.Error("FilesOnDir entry.Info(): %s ", err)
		}
		infos = append(infos, info)
	}

	return infos
}

func CompressGzip(path string) (string, error) {
	// Abrir archivo CSV original
	f, err := os.Open(path)
	defer f.Close()
	if err == nil {
		// Crear archivo comprimido
		g_path := path + ".gz"
		fw, err := os.Create(g_path)
		defer fw.Close()
		if err == nil {
			// Crear un escritor gzip
			gw := gzip.NewWriter(fw)
			defer gw.Close()

			// Crear un escritor CSV en el escritor gzip
			cw := csv.NewWriter(gw)

			// Leer el archivo original y escribir en el archivo comprimido
			records, err := csv.NewReader(f).ReadAll()
			if err == nil {
				err := cw.WriteAll(records)
				if err == nil {
					return g_path, nil
				} else {
					sLog.Error("compressGzip: cw.WriteAll(records): %s", err)
				}
			} else {
				sLog.Error("compressGzip: csv.NewReader(f).ReadAll(): %s", err)
			}
		} else {
			sLog.Error("compressGzip: os.Create(new_path): %s", err)
		}
	} else {
		sLog.Error("compressGzip: os.Open( %s ): %s", path, err)
	}
	return "", err
}

func DecompressGzip(path string) string {

	gzipfile, err := os.Open(path)
	if err == nil {

		reader, err := gzip.NewReader(gzipfile)
		if err == nil {
			defer reader.Close()

			new_path := strings.TrimSuffix(path, ".gz")

			writer, err := os.Create(new_path)
			if err == nil {
				defer writer.Close()

				_, err = io.Copy(writer, reader)
				if err == nil {
					return new_path
				} else {
					sLog.Error("decompressGzip: io.Copy(writer, reader): %s", err)
				}
			} else {
				sLog.Error("decompressGzip: os.Create( %s ): %s", new_path, err)
			}
		} else {
			sLog.Error("decompressGzip: gzip.NewReader( %s ): %s", path, err)
		}
	} else {
		sLog.Error("decompressGzip: os.Open( %s ): %s", path, err)
	}

	return ""
}

func DeletePath(path string) error {
	err := os.Remove(path)
	if err != nil {
		err = fmt.Errorf("os.Remove( %s ): %s", path, err)
	}
	return err
}

func CreateErrorDir(path string) (dir_path string, err error) {

	path_split := strings.Split(path, ".")

	dir_path = path_split[0]

	err = GetDir(dir_path)

	return
}

func GetDir(path string) (err error) {

	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			// File or directory does not exist
			err = os.MkdirAll(path, mode(0755, os.ModeDir))
			if err != nil {
				sLog.Error("GetDir: os.MkdirAll: %s", err)
			}
		} else {
			// Some other error. The file may or may not exist
			sLog.Error("GetDir: os.Stat( %q ): %s", path, err)
		}
	}

	return
}

// mode returns the file mode masked by the umask
func mode(mode, umask os.FileMode) os.FileMode {
	return mode & ^umask
}
