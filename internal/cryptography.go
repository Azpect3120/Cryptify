package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
)

// input: relative or absolute path of input file
//
// output: relative or absolute path of output file
//
// key: 32 byte encryption key
func EncryptFile (input string, output string, key []byte) error {
	// Define input path
	var inputPath string

	if IsAbsPath(input) {
		inputPath = input
	} else {
		absPath, err := GetAbsPath(input)
		if err != nil {
			return err
		}
		inputPath = absPath
	}

	// Define output path
	var outputPath string

	if IsAbsPath(output) {
		outputPath = output
	} else {
		absPath, err := GetAbsPath(output)
		if err != nil {
			return err
		}
		outputPath = absPath
	}

	// Convert input file to byte array
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// Create a new cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create byte array for encrypt file
	cipherText := make([]byte, aes.BlockSize + len(data))

	// Create initialization vector
	iv := cipherText[:aes.BlockSize]

	// Fill initialization vector with random values
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Create data encryptor stream
	encryptor := cipher.NewCFBEncrypter(block, iv)

	// Fill the encrypted array started AFTER the block with the newly encrypted data
	encryptor.XORKeyStream(cipherText[aes.BlockSize:], data)

	// Write encrypted data to the new file
	if err := os.WriteFile(outputPath, cipherText, 0644); err != nil {
		return err
	}

	return nil
}

// input: relative or absolute path of input file
//
// output: relative or absolute path of output file
//
// key: 32 byte encryption key
func DecryptFile (input string, output string, key []byte) error {
	// Define input path
	var inputPath string

	if IsAbsPath(input) {
		inputPath = input
	} else {
		absPath, err := GetAbsPath(input)
		if err != nil {
			return err
		}
		inputPath = absPath
	}

	// Define output path
	var outputPath string

	if IsAbsPath(output) {
		outputPath = output
	} else {
		absPath, err := GetAbsPath(output)
		if err != nil {
			return err
		}
		outputPath = absPath
	}

	// Create encrypted data and create data array
	ciphertext, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// Create new cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Ensure the length of the data array is larger than the block
	if len(ciphertext) < aes.BlockSize {
		return errors.New("Invalid key size for the encrypted file.")
	}

	// Create initialization vector
	iv := ciphertext[:aes.BlockSize]

	// Remove iv from the data array
	ciphertext = ciphertext[aes.BlockSize:]

	// Create data decryptor stream
	decryptor := cipher.NewCFBDecrypter(block, iv)

	// Decrypt data array using decryptor
	decryptor.XORKeyStream(ciphertext, ciphertext)

	// Write decrypted data into output file
	if err := os.WriteFile(outputPath, ciphertext, 0644); err != nil {
		return err
	}

	return nil
}

// fileName: name of the key file 
//
// filePath: directory path of the new key file
func CreateNewKeyFile (fileNam, filePath string) {

}
