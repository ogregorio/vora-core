package constants

import (
	errors "errors"
	files "graphgdb/utils/files"
	os "os"
	strings "strings"
)

var constants map[string]string

var instantiated bool = false

func Constants(path string) (bool, error) {

	dataConsts, err := files.Read(path)

	if err != nil {
		return false, errors.New("file not found")
	}

	constants = make(map[string]string, len(dataConsts))

	// for each const found in config file, split in two: a key (name of const) and a value
	for i := 0; i < len(dataConsts); i++ {

		c := strings.Split(dataConsts[i], " ")
		constants[c[0]] = c[1]

	}

	// now the dictionary of constants has been instantiated
	instantiated = true

	return true, nil
}

//GetConstant expected as first param a local to get const, and second param as value to find
func getConstant(param ...string) (string, error) {

	var concluded bool = true
	var err error

	// if dictionary of constants has not been instantiated
	if !instantiated {
		concluded, err = Constants(param[0])
	}

	// if constants was obtained with succesful
	// verify if constant requested exist, and return it
	if concluded {
		if val, ok := constants[param[1]]; ok {
			return val, nil
		}
	} else if !concluded && nil != err {
		return "", errors.New("file not found")
	} else if concluded && nil == err {
		return "", errors.New("constant " + param[1] + " not found")
	}

	return "", err
}

func GetConstant(param ...string) (string, error) {
	if len(param) == 1 {
		return getConstant(os.Getenv("GRAPHGDB_ROOT"), param[0])
	} else if len(param) == 2 {
		return getConstant(param[0], param[1])
	}
	return "", nil
}
