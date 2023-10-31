package main

import (
	"fmt"
	"os"

	"github.com/Azpect3120/Cryptify/internal"
)

func main() {
	priv, pub, err := internal.GenerateRSAKeys()
	if err != nil {
		fmt.Println("Fatal error: ", err)
		os.Exit(1)
	}

    fmt.Printf("%v\n\n", priv)
    fmt.Printf("%v\n\n", pub)

	internal.SavePrivateKeyToFile(priv, "private.pem")
	internal.SavePublicKeyToFile(pub, "public.pem")

	// Ensure valid usage
	if len(os.Args) < 2 {
		fmt.Println("Invalid arguments: use '-h' to get help.")
		os.Exit(1)
	}

	// Get flag from args
	flag := os.Args[1]

	// Main handler
	switch flag {
	// Encrypt flag
	case "-e":
		// Ensure valid usage
		if len(os.Args) < 4 {
			fmt.Println("Invalid usage: use '-h' to get help.")
			os.Exit(1)
		}

		// Define paths
		var keyPath string
		var inputPath string
		var outputPath string

		// Get key path from args
		path, err := internal.GetAbsPath(os.Args[2])
		if err != nil {
			fmt.Println("Fatal error: ", err.Error())
			os.Exit(1)
		}

		keyPath = path

		// Get input path from args
		path, err = internal.GetAbsPath(os.Args[3])
		if err != nil {
			fmt.Println("Fatal error: ", err.Error())
			os.Exit(1)
		}

		inputPath = path

		// Get output path if provided
		if len(os.Args) > 4 {
			path, err := internal.GetAbsPath(os.Args[4])
			if err != nil {
				fmt.Println("Fatal error: ", err.Error())
				os.Exit(1)
			}
			outputPath = path
		} else {
			outputPath = inputPath
		}

		// Encrypt file
		if err := internal.EncryptFile(inputPath, outputPath, keyPath); err != nil {
			fmt.Println("Fatal error: ", err.Error())
			os.Exit(1)
		} else {
			fmt.Println("File encrypted successfully.")
			os.Exit(0)
		}

	// Decrypt flag
	case "-d":
		// Ensure valid usage
		if len(os.Args) < 4 {
			fmt.Println("Invalid usage: use '-h' to get help.")
			os.Exit(1)
		}

		// Define paths
		var keyPath string
		var inputPath string
		var outputPath string

		// Get key path from args
		path, err := internal.GetAbsPath(os.Args[2])
		if err != nil {
			fmt.Println("Fatal error: ", err.Error())
			os.Exit(1)
		}

		keyPath = path

		// Get input path from args
		path, err = internal.GetAbsPath(os.Args[3])
		if err != nil {
			fmt.Println("Fatal error: ", err.Error())
			os.Exit(1)
		}

		inputPath = path

		// Get output path if provided
		if len(os.Args) > 4 {
			path, err := internal.GetAbsPath(os.Args[4])
			if err != nil {
				fmt.Println("Fatal error: ", err.Error())
				os.Exit(1)
			}
			outputPath = path
		} else {
			outputPath = inputPath
		}

		// Decrypt file
		if err := internal.DecryptFile(inputPath, outputPath, keyPath); err != nil {
			fmt.Println("Fatal error: ", err.Error())
			os.Exit(1)
		} else {
			fmt.Println("File decrypted successfully.")
			os.Exit(0)
		}

	// Key flags
	case "-k":
		// Ensure valid args were provided
		if len(os.Args) < 4 {
			fmt.Println("Invalid arguments: use '-h' to get help.")
			os.Exit(1)
		}

		// Define name and path
		var fileName string
		var dirPath string

		fileName = os.Args[2]
		dirPath = os.Args[3]

		// Create new key file
		if err := internal.CreateNewKeyFile(fileName, dirPath); err != nil {
			fmt.Println("Fatal error: ", err.Error())
			os.Exit(1)
		}

	// Help flag
	case "-h":
		fmt.Println("Encrypting a file:\ncryptify -e <key_path> <input_path> <?output_path>\n\nDecrypting a file:\ncryptify -d <key_path> <input_path> <?output_path>\n\nCreating a new encryption/decryption key file:\ncryptify -k <name> <location>")
		os.Exit(0)

	// Invalid usage
	default:
		fmt.Println("Invalid arguments: use '-h' to get help.")
		os.Exit(1)
	}
}
