package access

import (
	errors "errors"
	fmt "fmt"
	crypt "graphgdb/crypt"
	system "graphgdb/system"
	files "graphgdb/utils/files"

	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
)

var format string = "%s | %s"

//Access provides a first access function
func CreateAccess(login string, password string) error {
	data := fmt.Sprintf(format, login, password)
	return files.Write(
		viper.GetString("credentials"),
		crypt.Encrypt(crypt.Keygen(crypt.Generatrix()), data),
	)
}

//GrantAccess provides a method to grant permission to user
func GrantAccess(login string, password string) error {
	registeredUser, registeredPass := readRegisteredCredentials()
	if login == registeredUser && password == registeredPass {
		system.SetLogged(true)
		return nil
	} else {
		system.SetLogged(false)
		return errors.New("grant access failed, incorrect pass")
	}
}

//UngrantAccess provides a method to ungrant permission to user
func UngrantAccess() {
	system.SetLogged(false)
}

//VerifyAccess provides a access verify function
func VerifyAccess() bool {
	return system.IsLogged()
}

/*
	the encrypted file encrypts the credentials using an encrypted key
 	which is the result of the system root generatrix generated in setup
*/

//readAccess read from archive the credential data
func readRegisteredCredentials() (string, string) {
	content, err := files.Read(viper.GetString("credentials"))
	cobra.CheckErr(err)
	data := crypt.Decrypt(crypt.Keygen(crypt.Generatrix()), content[0])
	var user, password string
	fmt.Sscanf(data, format, &user, &password)
	return user, password
}
