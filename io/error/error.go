package io

import "fmt"

//FILE_NOT_FOUND occurs when a archive was not found
var FILE_NOT_FOUND int8 = 10

//ShowMessage show a lot of default error by a number passed in param
func ShowMessage(value int8) {
	switch value {
	case FILE_NOT_FOUND:
		fmt.Printf("File not found!")
	}
}
