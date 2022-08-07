package main

import (
	"fmt"
	"github.com/gsoultan/go-cipher/cipher"
)

func main() {
	text := "hello worldssssss"
	key := "12345678910111213141516171819201"

	cipher := cipher.NewGoCipher(key)
	encryptedText, err := cipher.Encrypt(text)
	fmt.Println(encryptedText)

	pt, err := cipher.Decrypt(encryptedText)
	fmt.Println(err)
	fmt.Println(pt)
}
