package access

import (
	errors "errors"
	crypt "graphgdb/crypt"
	errorTreat "graphgdb/io/error"
	files "graphgdb/utils/files"
	"strings"

	viper "github.com/spf13/viper"
)

type Credentials struct {
	login    string
	password string
	sha3     string
}

var credentials Credentials = Credentials{login: "", password: ""}

func Access(login string, password string) {
	credentials = Credentials{login: login, password: password}
	credentials.sha3 = string(crypt.Keygen(login, password))
	burnAccess()
}

func GetKeygenLocal() string {
	readAccess()
	return credentials.sha3
}

func VerifyAccess() (bool, error) {
	if credentials.sha3 == string(crypt.Keygen(credentials.login, credentials.password)) {
		return true, nil
	} else {
		return false, errors.New("unverified access")
	}
}

func burnAccess() {
	err := files.WriteAllBytes(
		viper.GetString("credential"),
		[]byte(crypt.Keygen(getCredentials(), getGeneratrix())),
	)
	errorTreat.CheckError(err)
}

func readAccess() {
	data, err := files.Read(viper.GetString("credential"))
	errorTreat.CheckError(err)
	access := crypt.Decrypt(
		crypt.Keygen(getCredentials(), getGeneratrix()),
		data[0])
	values := strings.Split(string(access), " ")
	credentials = Credentials{login: values[0], password: values[1], sha3: values[2]}
}

func getCredentials() string {
	return credentials.login + " " + credentials.password + " " + credentials.sha3
}

func getGeneratrix() string {
	root, err := files.Read(viper.GetString("root"))
	errorTreat.CheckError(err)
	return root[0]
}
