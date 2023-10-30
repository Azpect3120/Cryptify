package main

import (
	"fmt"
	"os"

	"github.com/Azpect3120/Cryptify/internal"
)

// Define base string
// WIP: BUILD FILES FOR KEYS
const key string = "q6FlbwZfFUw4fvlRiQReIxxn6sWrQDSQ"

func main () {
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
            if len(os.Args) < 3 {
                fmt.Println("Usage: program <flag> <path>")
                os.Exit(1)
            }

            // Define paths
            var inputPath string
            var outputPath string
            
            // Get input path from args
            path, err := internal.GetAbsPath(os.Args[2])
            if err != nil {
                fmt.Println("Fatal error: ", err.Error())
                os.Exit(1)
            }
            
            inputPath = path

            // Get output path if provided
            if len(os.Args) > 3 {
                path, err := internal.GetAbsPath(os.Args[3])
                if err != nil {
                    fmt.Println("Fatal error: ", err.Error())
                    os.Exit(1)
                }
                outputPath = path
            } else {
                outputPath = inputPath
            }

            // Encrypt file
            if err := internal.EncryptFile(inputPath, outputPath, []byte(key)); err != nil {
                fmt.Println("Fatal error: ", err.Error())
                os.Exit(1)
            } else {
                fmt.Println("File encrypted successfully.")
                os.Exit(0)
            }

        // Decrypt flag
        case "-d":
            // Ensure valid usage
            if len(os.Args) < 3 {
                fmt.Println("Usage: program <flag> <path>")
                os.Exit(1)
            }

            // Define paths
            var inputPath string
            var outputPath string
            
            // Get input path from args
            path, err := internal.GetAbsPath(os.Args[2])
            if err != nil {
                fmt.Println("Fatal error: ", err.Error())
                os.Exit(1)
            }
            
            inputPath = path

            // Get output path if provided
            if len(os.Args) > 3 {
                path, err := internal.GetAbsPath(os.Args[3])
                if err != nil {
                    fmt.Println("Fatal error: ", err.Error())
                    os.Exit(1)
                }
                outputPath = path
            } else {
                outputPath = inputPath
            }

            // Decrypt file
            if err := internal.DecryptFile(inputPath, outputPath, []byte(key)); err != nil {
                fmt.Println("Fatal error: ", err.Error())
                os.Exit(1)
            } else {
                fmt.Println("File decrypted successfully.")
                os.Exit(0)
            }

        // Key flags
        case "-k":
            // Ensure valid args were provided
            if len(os.Args) < 5 {
                fmt.Println("Invalid arguments: use '-h' to get help.")
                os.Exit(1)
            }

            // Ensure new keyword is used
            if os.Args[2] != "new" {
                fmt.Println("Invalid argument '"+ os.Args[2] + "': use '-h' to get help.")
                os.Exit(1)
            } 

            // Define name and path
            var fileName string
            var dirPath string

            fileName = os.Args[3]
            dirPath = os.Args[4]

            // Create new key file
            internal.CreateNewKeyFile(fileName, dirPath)

        // Help flag
        case "-h":
            fmt.Println("This is the help menu.")
            os.Exit(0)

        // Invalid usage
        default:
            fmt.Println("Invalid arguments: use '-h' to get help.")
            os.Exit(1)
    }
}
