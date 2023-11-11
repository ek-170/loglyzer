package util

import (
	"errors"
	"log"
	"os"
)

// this method ignore subdirectory
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

func SearchFile(dir string, q string) (string, error) {
    res := ""
    files, err := ScanFiles(dir)
    if err != nil {
        return res, err
    }
    if len(files) == 0 {
        return res, errors.New("there is no file you want to parse")
    }
    for _, f := range files {
        if f == q {
            res = f
        }
    }
    return res, nil
}