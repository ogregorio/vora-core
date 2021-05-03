package crypt

import (
	aes "crypto/aes"
	hex "encoding/hex"
	errorTreat "graphgdb/io/error"
	files "graphgdb/utils/files"

	viper "github.com/spf13/viper"
	crypto "golang.org/x/crypto/sha3"
)

//Encrypt expects an array of bytes as the key and a data string as the parameter, and return an encoded string.
func Encrypt(key []byte, data string) string {
	cypher, err := aes.NewCipher(key)
	errorTreat.CheckError(err)

	// allocate space for ciphered data
	encrypted := make([]byte, len(data))

	// encrypt
	cypher.Encrypt(encrypted, []byte(data))
	// return hex string
	return hex.EncodeToString(encrypted)
}

//Decrypt expects an array of bytes as the key and a data string as the parameter, and return an decoded string.
func Decrypt(key []byte, data string) string {
	ciphertext, _ := hex.DecodeString(data)

	cypher, err := aes.NewCipher(key)
	errorTreat.CheckError(err)

	// allocate space for decrypted data
	decrypted := make([]byte, len(ciphertext))
	cypher.Decrypt(decrypted, ciphertext)

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
