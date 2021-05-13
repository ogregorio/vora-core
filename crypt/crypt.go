package crypt

import (
	errorTreat "vora/io/error"
	files "vora/utils/files"

	openssl "github.com/Luzifer/go-openssl/v4"
	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
	crypto "golang.org/x/crypto/sha3"
)

var o *openssl.OpenSSL = openssl.New()

//Encrypt expects an array of bytes as the key and a data string as the parameter, and return an encoded string.
func Encrypt(key []byte, data string) string {
	encrypted, err := o.EncryptBytes(string(key), []byte(data), openssl.BytesToKeySHA256)
	cobra.CheckErr(err)
	return string(encrypted)
}

//Decrypt expects an array of bytes as the key and a data string as the parameter, and return an decoded string.
func Decrypt(key []byte, data string) string {
	decrypted, err := o.DecryptBytes(string(key), []byte(data), openssl.BytesToKeySHA256)
	cobra.CheckErr(err)
	return string(decrypted)
}

//KeygenStr expects a string as param and strings to concat, and return an key
func Keygen(param ...string) []byte {
	hasher := crypto.New256()
	var content string
	for i := 0; i < len(param); i++ {
		content += param[i]
	}
	hasher.Write([]byte(content))
	return hasher.Sum(nil)
}

//Generatrix delivers an extra layer for encryption based on system data
func Generatrix() string {
	root, err := files.Read(viper.GetString("root"))
	errorTreat.CheckError(err)
	return root[0]
}
