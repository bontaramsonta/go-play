package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
)

func keyToCipher(key string) (cipher.Block, error) {
	return aes.NewCipher([]byte(key))
}

func DebugEncryptDecrypt(masterKey, iv, password string) (string, string) {
	encryptedPassword := encrypt(password, masterKey, iv)
	decryptedPassword := decrypt(encryptedPassword, masterKey, iv)
	return encryptedPassword, decryptedPassword
}

// don't touch below this line

func encrypt(plainText, key, iv string) string {
	bytes := []byte(plainText)
	blockCipher, err := keyToCipher(key)
	if err != nil {
		log.Println(err)
		return ""
	}
	stream := cipher.NewCTR(blockCipher, []byte(iv))
	stream.XORKeyStream(bytes, bytes)
	return fmt.Sprintf("%x", bytes)
}

func decrypt(cipherText, key, iv string) string {
	blockCipher, err := keyToCipher(key)
	if err != nil {
		log.Println(err)
		return ""
	}
	stream := cipher.NewCTR(blockCipher, []byte(iv))
	bytes, err := hex.DecodeString(cipherText)
	if err != nil {
		log.Println(err)
		return ""
	}
	stream.XORKeyStream(bytes, bytes)
	return string(bytes)
}
