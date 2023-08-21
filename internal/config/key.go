package config

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/dianhadi/user/pkg/log"
)

var PublicKey *rsa.PublicKey

func LoadPublicKey(filename string) error {
	// Read public key from file
	keyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("Error reading public key:", err)
		return err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		log.Error("Error parsing public key: ", err)
		return err
	}
	PublicKey = key
	return nil
}
