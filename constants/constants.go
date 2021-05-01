package constants

import (
	"errors"
	files "graphgdb/utils/files"
	"path/filepath"
	"strings"
)

var constants map[string]string

var instantiated bool = false

func Constants() (bool, error) {
	absolutePath, _ := filepath.Abs("./config/.config")
	dataConsts, err := files.Read(absolutePath)

	if err != nil {
		return false, errors.New("file not found")
	} //if error was found

	constants = make(map[string]string, len(dataConsts))

	for i := 0; i < len(dataConsts); i++ {

		c := strings.Split(dataConsts[i], " ")
		constants[c[0]] = c[1]

	}

	instantiated = true

	return true, nil
}

func GetConstant(name string) (string, error) {

	var concluded bool
	var err error

	if !instantiated {
		concluded, err = Constants()
	}

	if concluded {
		for i := 0; i < len(constants); i++ {
			if val, ok := constants[name]; ok {
				return val, nil
			}
		}
	} else if !concluded && nil != err {
		return "", errors.New("file not found")
	} else if concluded && nil == err {
		return "", errors.New("constant " + name + " not found")
	}

	return "", err
}
