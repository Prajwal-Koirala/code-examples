package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

func main() {
	// Generate a private and a public key using rsa.
	privateKey, publicKey := generatePrivateAndPublicRSAKey()
	fmt.Printf("Private key: %s\n", privateKey)
	fmt.Printf("Public key: %s\n", publicKey)
}

// Generate a private and a public key using rsa.
func generatePrivateAndPublicRSAKey() ([]byte, []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatalln(err)
	}
	return privateKey.D.Bytes(), privateKey.PublicKey.N.Bytes()
}
