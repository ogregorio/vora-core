package crypt

import (
	aes "crypto/aes"
	hex "encoding/hex"
	error "graphgdb/io/error"

	crypto "golang.org/x/crypto/sha3"
)

//Encrypt expects an array of bytes as the key and a data string as the parameter, and return an encoded string.
func Encrypt(key []byte, data string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	error.CheckError(err)

	// allocate space for ciphered data
	out := make([]byte, len(data))

	// encrypt
	c.Encrypt(out, []byte(data))
	// return hex string
	return hex.EncodeToString(out)
}

//Decrypt expects an array of bytes as the key and a data string as the parameter, and return an decoded string.
func Decrypt(key []byte, data string) string {
	ciphertext, _ := hex.DecodeString(data)

	c, err := aes.NewCipher(key)
	error.CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s
}

//KeygenStr expects a string as param and strings to concat, and return an key
func Keygen(param ...string) []byte {
	hasher := crypto.New512()
	var content string
	for i := 0; i < len(param); i++ {
		content += param[i]
	}
	hasher.Write([]byte(content))
	return hasher.Sum(nil)
}
