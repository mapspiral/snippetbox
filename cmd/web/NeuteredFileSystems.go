package main

import (
	"net/http"
	"path/filepath"
)

type neuteredFileSystem struct {
	fileSystem http.FileSystem
}

func (target neuteredFileSystem) Open(path string) (http.File, error) {
	file, error := target.fileSystem.Open(path)
	if error != nil {
		return nil, error
	}

	fileInfo, error := file.Stat()
	if error != nil {
		return nil, error
	}

	if fileInfo.IsDir() {
		indexFileName := filepath.Join(path, "index.html")
		if _, error := target.Open(indexFileName); error != nil {
			closeError := file.Close()
			if closeError != nil {
				return nil, closeError
			}
			return nil, error
		}
	}

	return file, nil
}
