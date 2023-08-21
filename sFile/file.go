package sFile

import (
	"archive/zip"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/yasseldg/simplego/sLog"
	"github.com/yasseldg/simplego/sStr"
)

func FilesOnDir(dir_path string, patterns ...string) []fs.FileInfo {
	entries, err := os.ReadDir(dir_path)
	if err != nil {
		sLog.Error("FilesOnDir os.ReadDir( %s ): %s", dir_path, err)
		return nil
	}

	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			sLog.Error("FilesOnDir entry.Info(): %s ", err)
			continue
		}
		if sStr.FindPatterns(info.Name(), patterns...) {
			infos = append(infos, info)
		}
	}
	return infos
}

func CompressGzip(file_path string) (string, error) {
	// Abrir archivo CSV original
	f, err := os.Open(file_path)
	if err != nil {
		sLog.Error("CompressGzip: os.Open( %s ): %s", file_path, err)
		return "", err
	}
	defer f.Close()

	// Crear archivo comprimido
	gz_path := file_path + ".gz"
	fw, err := os.Create(gz_path)
	if err != nil {
		sLog.Error("CompressGzip: os.Create(new_path): %s", err)
		return "", err
	}
	defer fw.Close()

	// Crear un escritor gzip
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// Crear un escritor CSV en el escritor gzip
	cw := csv.NewWriter(gw)

	// Leer el archivo original y escribir en el archivo comprimido
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		sLog.Error("CompressGzip: csv.NewReader(f).ReadAll(): %s", err)
		return "", err
	}

	err = cw.WriteAll(records)
	if err != nil {
		sLog.Error("CompressGzip: cw.WriteAll(records): %s", err)
		return "", err
	}
	return gz_path, nil
}

func DecompressGzip(file_path string) string {
	gzipfile, err := os.Open(file_path)
	if err != nil {
		sLog.Error("DecompressGzip: os.Open( %s ): %s", file_path, err)
		return ""
	}
	defer gzipfile.Close()

	reader, err := gzip.NewReader(gzipfile)
	if err != nil {
		sLog.Error("DecompressGzip: gzip.NewReader( %s ): %s", file_path, err)
		return ""
	}
	defer reader.Close()

	new_path := strings.TrimSuffix(file_path, ".gz")
	writer, err := os.Create(new_path)
	if err != nil {
		sLog.Error("DecompressGzip: os.Create( %s ): %s", new_path, err)
		return ""
	}
	defer writer.Close()

	_, err = io.Copy(writer, reader)
	if err != nil {
		sLog.Error("DecompressGzip: io.Copy(writer, reader): %s", err)
		return ""
	}

	return new_path
}

func DecompressZip(file_path, dir_path string) error {
	reader, err := zip.OpenReader(file_path)
	if err != nil {
		sLog.Error("DecompressZip: zip.OpenReader( %s ): %s", file_path, err)
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		err := decompressZipFile(file, dir_path)
		if err != nil {
			sLog.Error("DecompressZip: decompressZipFile( %s ): %s", file.Name, err)
			return err
		}
	}
	return nil
}

func decompressZipFile(file *zip.File, dir_path string) error {
	filePath := filepath.Join(dir_path, file.Name)

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		sLog.Error("decompressZipFile: os.MkdirAll(filepath.Dir( %s ): %s", filePath, err)
		return err
	}

	if file.FileInfo().IsDir() {
		return nil
	}

	outputFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		sLog.Error("decompressZipFile: os.OpenFile( %s ): %s", filePath, err)
		return err
	}
	defer outputFile.Close()

	zipFile, err := file.Open()
	if err != nil {
		sLog.Error("decompressZipFile: file.Open(): %s", err)
		return err
	}
	defer zipFile.Close()

	_, err = io.Copy(outputFile, zipFile)
	if err != nil {
		sLog.Error("decompressZipFile: io.Copy(outputFile, zipFile): %s", err)
		return err
	}
	return nil
}

func DeletePath(file_path string) error {
	exist, err := ExistingPath(file_path)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}

	err = os.Remove(file_path)
	if err != nil {
		err = fmt.Errorf("os.Remove( %s ): %s", file_path, err)
	}
	return err
}

func ExistingPath(file_path string) (bool, error) {
	_, err := os.Stat(file_path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateErrorDir(file_path string) (dir_path string, err error) {

	path_split := strings.Split(file_path, ".")

	dir_path = path_split[0]

	err = GetDir(dir_path)

	return
}

func GetDir(file_path string) (err error) {
	_, err = os.Stat(file_path)
	if err == nil {
		// Directory exists
		return nil
	}

	if os.IsNotExist(err) {
		// File or directory does not exist
		return os.MkdirAll(file_path, mode(0755, os.ModeDir))
	}

	sLog.Error("GetDir: os.Stat( %q ): %s", file_path, err)
	return err
}

// mode returns the file mode masked by the umask
func mode(mode, umask os.FileMode) os.FileMode {
	return mode & ^umask
}
