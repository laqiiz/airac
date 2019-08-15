package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"time"
)

// Private key
var privateKey *rsa.PrivateKey

// Verify key
var verifyKey *rsa.PublicKey

func init() {
	// Load private key
	privatePath := os.Getenv("PRIVATE_KEY_PATH")
	if privatePath == "" {
		panic("env PRIVATE_KEY_PATH is required")
	}

	privateBytes, err := ioutil.ReadFile(privatePath)
	if err != nil {
		panic(err)
	}
	loadPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		panic(err)
	}
	privateKey = loadPrivateKey

	// Load public key
	publicPath := os.Getenv("PUBLIC_KEY_PATH")
	if publicPath == "" {
		panic("env PUBLIC_KEY_PATH is required")
	}

	publicBytes, err := ioutil.ReadFile(publicPath)
	if err != nil {
		panic(err)
	}

	loadVerifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		panic(err)
	}

	verifyKey = loadVerifyKey
}

func GenerateToken(id string) (string, error) {
	// create token
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // TODO Magic Number
		Id:        "",
		IssuedAt:  0,
		Issuer:    id, // TODO 適切な値
		NotBefore: 0,
		Subject:   "test", // test
	})

	// signed
	return tkn.SignedString(privateKey)
}

// Refs: https://github.com/dgrijalva/jwt-go/blob/master/rsa_test.go#L73
func Verify(tkn *jwt.Token) error {
	method, castErr := tkn.Method.(*jwt.SigningMethodRSA)
	if !castErr {
		return fmt.Errorf("unexpected signing method: %v", tkn.Header["alg"])
	}

	return method.Verify(tkn.Signature, tkn.Method.Alg(), verifyKey)
}
