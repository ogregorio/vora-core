package access

import (
	constants "graphgdb/constants"
	error "graphgdb/io/error"
	files "graphgdb/utils/files"
)

var KEY string = ""
var keyFile string

func Access(login string, password string) bool {
	KEY := Login(login, password)
	keyFile, err := constants.GetConstant("KEYS_FILE")
	error.CheckError(err)
	files.Write(keyFile, KEY)
	return KEY != ""
}

func GetAccess() string {
	access, err := files.Read(keyFile)
	error.CheckError(err)
	return access[0]
}
