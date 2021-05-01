package constants

import (
	"errors"
	"strconv"
)

func ConstantsTest() (bool, error) {
	for i := 0; i < 10; i++ {
		toPrint, err := GetConstant("./config/.configtest", "TEST_PATH_"+strconv.Itoa(i))

		if toPrint != "PASS" {
			return false, errors.New("error in file read")
		} else if err != nil {
			return false, errors.New("error in constants test")
		}

	}

	return true, nil
}
