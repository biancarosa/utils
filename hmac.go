package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
)

func main() {
	var secretKey = flag.String("secretKey", "", "secret key to encrypt")
	var text = flag.String("text", "", "text to encrypt with secret key")

	flag.Parse()

	fmt.Printf("Secret Key :: %s\n", *secretKey)
	fmt.Printf("Text :: %s\n", *text)

	sig := hmac.New(sha512.New, []byte(*secretKey))
	sig.Write([]byte(*text))

	fmt.Println(hex.EncodeToString(sig.Sum(nil)))
}
