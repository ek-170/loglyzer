package util

import (
	"errors"
	"log"
	"os"
)

func ScanFiles(dir string) ([]string, error) {
    files, err := os.ReadDir(dir)
    if err != nil {
        log.Print("can't scan log directory")
        return nil, errors.New("can't scan log directory")
    }
    paths := []string{}
    for _, file := range files {
        if !file.IsDir() {
            paths = append(paths, file.Name())
        }
    }
    return paths, nil
}