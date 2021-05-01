package io

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

//Read get from a archive a vector of strings
func Read(path string) ([]string, error) {

	archive, err := os.Open(absolutePath(path)) //open file

	if err != nil {
		return nil, err
	} //if error was found

	defer archive.Close() //archive will be closed

	// Read file with scanner
	var lines []string
	scanner := bufio.NewScanner(archive)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err() //return lines or error
}

//WriteMultiLine write in a file a lot of lines in a string vector
func Write(path string, content ...string) error {

	archive, err := os.Create(absolutePath(path)) //create archive to write

	if err != nil {
		return err
	} //if error was found

	defer archive.Close() //archive will be closed

	// Write file with Writer
	writer := bufio.NewWriter(archive)
	for _, content := range content {
		fmt.Fprintln(writer, content)
	}

	return writer.Flush() //return lines or error
}

func absolutePath(path string) string {
	absolutePath, _ := filepath.Abs("./config/.config")

	// if has a path param, change default source
	if path != "" {
		absolutePath, _ = filepath.Abs(path)
	}

	return absolutePath
}