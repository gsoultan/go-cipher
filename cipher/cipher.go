package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/gsoultan/go-cipher/base64"
	"io"
)

type Cipher interface {
	Encrypt(plainText string) (encryptedText string, err error)
	Decrypt(encryptedText string) (plainText string, err error)
}

type goChiper struct {
	key string
}

func (g *goChiper) Encrypt(plainText string) (encryptedText string, err error) {
	if len(g.key) != 32 {
		return "", errors.New("invalid key size. key must be 32 character")
	}
	var a cipher.Block
	if a, err = aes.NewCipher([]byte(g.key)); err != nil {
		return "", err
	}

	var gcm cipher.AEAD
	if gcm, err = cipher.NewGCM(a); err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	e := gcm.Seal(nonce, nonce, []byte(plainText), []byte(g.key))
	return base64.EncodeBase64(e), nil
}

func (g *goChiper) Decrypt(encryptedText string) (plainText string, err error) {
	if len(g.key) != 32 {
		return "", errors.New("invalid key size. key must be 32 character")
	}

	var a cipher.Block
	if a, err = aes.NewCipher([]byte(g.key)); err != nil {
		return "", err
	}

	var gcm cipher.AEAD
	if gcm, err = cipher.NewGCM(a); err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedText) < nonceSize {
		return "", errors.New("invalid size")
	}

	var et, pt []byte

	if et, err = base64.DecodeBase64(encryptedText); err != nil {
		return "", err
	}
	nonce, cipherText := et[:nonceSize], et[nonceSize:]
	if pt, err = gcm.Open(nil, nonce, cipherText, []byte(g.key)); err != nil {
		return "", err
	}

	return string(pt[:]), nil

}

func NewGoCipher(key string) Cipher {
	a := new(goChiper)
	a.key = key
	return a
}
