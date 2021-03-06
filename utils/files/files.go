package io

import (
	bufio "bufio"
	fmt "fmt"
	ioutil "io/ioutil"
	os "os"
)

//Read get from a archive a vector of strings
func Read(path string) ([]string, error) {

	archive, err := os.Open(path) //open file

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

	return lines, scanner.Err() //return lines and error
}

//WriteMultiLine write in a file a lot of lines in a string vector
func Write(path string, content ...string) error {

	archive, err := os.Create(path) //create archive to write

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

//ReadAllBytes get from a archive a bytes's array
func ReadAllBytes(path string) ([]byte, error) {

	archive, err := os.Open(path) //open file

	if err != nil {
		return nil, err
	} //if error was found

	content, err := ioutil.ReadAll(archive)

	defer archive.Close() //archive will be closed

	return content, err //return bytes and error
}

//WriteAllBytes write in a archive a []byte
func WriteAllBytes(path string, content []byte) error {

	archive, err := os.Create(path) //create archive to write

	if err != nil {
		return err
	} //if error was found

	defer archive.Close() //archive will be closed

	// Write file with archive
	writedBytes, err := archive.Write(content)
	if writedBytes > 0 && err != nil {
		return nil
	} else {
		return err
	}
}
