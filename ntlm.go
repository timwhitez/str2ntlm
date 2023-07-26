package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/md4"
	"os"
	"strings"
	"unicode/utf16"
)

func CreateHash(str string) string {
	hasher := md4.New()
	hasher.Write(encodePassword(str))
	return strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
}

func encodePassword(password string) []byte {
	encoded := utf16.Encode([]rune(password))

	buff := make([]byte, len(encoded)*2)
	for i := 0; i < len(encoded); i++ {
		buff[i*2] = byte(encoded[i])
		buff[i*2+1] = byte(encoded[i] >> 8)
	}

	return buff
}

func main() {
	s := os.Args[1]
	hash := CreateHash(s)
	fmt.Println("NTLM hash:", hash)
}
