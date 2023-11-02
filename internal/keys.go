package internal

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// Generate new blank public/private key pairs
func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func SavePrivateKeyToFile (privateKey *rsa.PrivateKey, filename string) error {
	keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: keyBytes}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	err = pem.Encode(file, pemBlock)
	return err
}

func SavePublicKeyToFile (publicKey *rsa.PublicKey, filename string) error {
	keyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	pemBlock := &pem.Block{Type: "RSA KEY", Bytes: keyBytes}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	err = pem.Encode(file, pemBlock)
	return err
}

// Convert an RSA private key to an AES encryption/decryption key
func DeriveAESKeyFromRSAKey (privateKey *rsa.PrivateKey) ([]byte, error) {
	privateKeyDER, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(privateKeyDER)
	return hash[:], err
}


func LoadPrivateKey (filename string) (*rsa.PrivateKey, error){
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func LoadPublicKey (filename string) (*rsa.PublicKey, error){
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, err
	}

	return rsaPublicKey, nil
}
