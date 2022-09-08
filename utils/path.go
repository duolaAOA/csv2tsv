package utils

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

func ListFiles(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if IsCsvFile(path) {
			*files = append(*files, path)
		}
		return nil
	}
}

func PathExists(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsCsvFile(filename string) bool {
	suffix := path.Ext(filename)
	if suffix != ".csv" {
		return false
	}
	return true
}
