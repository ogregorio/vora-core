package access

import (
	"encoding/hex"

	crypto "golang.org/x/crypto/sha3"
)

func Login(login string, password string) string {
	hasher := crypto.New512()
	hasher.Write([]byte(login + password))
	return hex.EncodeToString(hasher.Sum(nil))
}
