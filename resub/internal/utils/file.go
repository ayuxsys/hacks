package utils

import (
	"bufio"
	"io"
	"os"

	"log"
)

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func FilePathToSlice(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, WrapError("os.Open", err)
	}
	defer file.Close()
	return FileToSlice(file), nil
}

func FileToSlice(file io.Reader) []string {
	var slice []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err().Error())
	}
	return slice
}
