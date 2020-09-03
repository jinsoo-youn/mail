package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

type AESCipher struct {
	block cipher.Block
}

func NewAesCipher(key []byte) (*AESCipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &AESCipher{block}, nil
}

func (a *AESCipher) EncryptString(s string) string {
	byteString := []byte(s)
	encryptByteArray := make([]byte, aes.BlockSize+len(s))

	iv := encryptByteArray[:aes.BlockSize]

	io.ReadFull(rand.Reader, iv)

	stream := cipher.NewCFBEncrypter(a.block, iv)
	stream.XORKeyStream(encryptByteArray[aes.BlockSize:], byteString)

	return base64.URLEncoding.EncodeToString(encryptByteArray)
}

func (a *AESCipher) DecryptString(base64String string) string {

	b, _ := base64.URLEncoding.DecodeString(base64String)
	byteString := []byte(b)

	decryptByteArray := make([]byte, len(byteString))
	iv := byteString[:aes.BlockSize]

	stream := cipher.NewCFBDecrypter(a.block, iv)
	stream.XORKeyStream(decryptByteArray, byteString[aes.BlockSize:])

	return string(decryptByteArray)
}
