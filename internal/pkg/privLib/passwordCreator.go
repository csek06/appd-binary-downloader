package privlib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func encrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}
func decrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

/*
PasswordCreator will encrypt a new password giving a base64 version of the ciphertext:key
*/
func PasswordCreator(password string) string {
	data := []byte(password)
	key, err := generateKey()
	if err != nil {
		log.Fatal(err)
	}
	ciphertext, err := encrypt(key, data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("ciphertext: %s\n", hex.EncodeToString(ciphertext))

	secretbyte := append(ciphertext[:], []byte(":")...)
	secretbyte = append(secretbyte[:], key[:]...)
	// cipher:key
	secret64 := base64.StdEncoding.EncodeToString(secretbyte)

	return secret64
}

/*
PasswordDecryptor will accept a base64 string of ciphertext:key and decrypt to plain text
*/
func PasswordDecryptor(secret64 string) string {
	//fmt.Println("Decrypting...")

	unsecret64, err := base64.StdEncoding.DecodeString(secret64)
	if err != nil {
		fmt.Println(err)
	}
	vals := strings.Split(string(unsecret64), ":")

	key := []byte(vals[1])
	cipher := []byte(vals[0])

	plaintext, err := decrypt(key, cipher)
	if err != nil {
		log.Fatal(err)
	}

	return string(plaintext)

}
