package access

import (
	errors "errors"
	"fmt"
	crypt "graphgdb/crypt"
	errorTreat "graphgdb/io/error"
	files "graphgdb/utils/files"
	strings "strings"

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

func VerifyAccess() (bool, error) {
	accessKey, err := readAccess()
	if credentials.sha3 == accessKey {
		return true, nil
	} else {
		return false, err
	}
}

/*
	the encrypted file encrypts the credentials using an encrypted key
 	(which is the result of the system root generatrix generated in setup + the credential itself)
*/
func burnAccess() {
	encrypted := crypt.Encrypt(
		[]byte(
			crypt.Keygen(
				crypt.Generatrix(),
			),
		), getCredentials(),
	)
	err := files.Write(viper.GetString("credential"), encrypted)
	errorTreat.CheckError(err)
}

func readAccess() (string, error) {
	data, err := files.Read(viper.GetString("credential"))
	errorTreat.CheckError(err)
	decrypted := crypt.Decrypt(
		[]byte(
			crypt.Keygen(
				crypt.Generatrix(),
			),
		), data[0],
	)
	access := strings.Split(string(decrypted), " ")
	if len(access) == 3 {
		credentials = Credentials{login: access[0], password: access[1], sha3: access[2]}
		return access[2], nil
	} else {
		fmt.Println(credentials)
		return "err", errors.New("impossible to read access")
	}
}

func getCredentials() string {
	return credentials.login + " " + credentials.password + " " + credentials.sha3
}
