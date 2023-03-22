package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	math_rand "math/rand"
	"time"
)

// var key = "a33c217a8db60a7ef3a7360f76a3157b7641315f7b80881b2cced8447a4c516ed40983ba662271a18af5742cece4c735b1bc45661fa95bb7493b00c5"
// var key = config.Env.ENCRYPTION_KEY
var key = "kN6ulzxbp29cKMGTw5lMyLMzdz7jkqn5"
var Rand *math_rand.Rand = math_rand.New(math_rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EncryptParamsValue(param string) string {
	k, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return ""
	}
	dataEncryption, _ := EncryptAES256GCMHex(param, k)
	if err != nil {
		return ""
	}
	return dataEncryption
}
func EncryptParamsValueConfig(param string, keyy string) string {
	k, err := base64.StdEncoding.DecodeString(keyy)
	if err != nil {
		return ""
	}
	dataEncryption, _ := EncryptAES256GCMHex(param, k)
	if err != nil {
		return ""
	}
	return dataEncryption
}

func DecryptParamsValue(encryptString string) string {
	k, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return ""
	}
	dataDecryption, err := DecryptAES256GCMHex(encryptString, k)
	if err != nil {
		return ""
	}
	return dataDecryption
}

func DecryptParamsValueConfig(encryptString string, keyy string) string {
	k, err := base64.StdEncoding.DecodeString(keyy)
	if err != nil {
		return ""
	}
	dataDecryption, err := DecryptAES256GCMHex(encryptString, k)
	if err != nil {
		return ""
	}
	return dataDecryption
}

// EncryptAES256GCMHex is Encrypt AES256 GCM
func EncryptAES256GCMHex(data string, key []byte) (string, error) {
	text := []byte(data)

	if len(key) == 0 {
		errMsg := "Error Get Config Key"
		return "", fmt.Errorf("%v", errMsg)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	seal := gcm.Seal(nil, nonce, text, nil)

	output := fmt.Sprintf("%x%x", nonce, seal)
	return output, nil
}

// DecryptAES256GCMHex is Encrypt AES256 GCM
func DecryptAES256GCMHex(encrypt string, key []byte) (string, error) {
	ciphertext, err := hex.DecodeString(encrypt)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[Rand.Intn(len(charset))]
	}
	return string(result)
}

func PrintEncryptAES256GCMHex(encrypt string) {

	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("************************     Encrypt base64      **************************")
	fmt.Println("************************   EncryptAES256GCMHex   **************************")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println(string(encrypt))
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("***************************************************************************")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println()

}

func PrintDecryptAES256GCMHex(decrypt string) {

	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("************************     Decrypt base64      **************************")
	fmt.Println("************************   DecryptAES256GCMHex   **************************")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println(string(decrypt))
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("***************************************************************************")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println()

}
