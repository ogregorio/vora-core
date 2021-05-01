package constants

import (
	"errors"
	"strconv"
)

func ConstantsTest() (bool, error) {
	for i := 0; i < 10; i++ {
		toPrint, err := GetConstant("TEST_PATH_" + strconv.Itoa(i))
		if toPrint == "PASS" && err == nil {
			continue
		} else if nil != err {
			return false, errors.New("error in file read")
		} else {
			return false, errors.New("error in constants test")
		}
	}

	return true, nil
}
