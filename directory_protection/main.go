package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type AppContext struct {
	PasswordHash      string
	ProtectedNames    []string
	ProtectionEnabled bool
	ProtectedFile     string
}

func NewAppContext() *AppContext {
	return &AppContext{
		ProtectionEnabled: false,
		ProtectedFile:     "template.tpl",
	}
}

func (ctx *AppContext) loadProhibitedNames() error {
	data, err := ioutil.ReadFile(ctx.ProtectedFile)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	ctx.PasswordHash = lines[0]
	ctx.ProtectedNames = lines[1:]
	return nil
}

func (ctx *AppContext) toggleProtection() {
	ctx.ProtectionEnabled = !ctx.ProtectionEnabled
	var permissions os.FileMode
	if ctx.ProtectionEnabled {
		fmt.Println("Protection mode is now enabled.")
		permissions = 0000
	} else {
		fmt.Println("Protection mode is now disabled.")
		permissions = 0777
	}

	for _, fileName := range ctx.ProtectedNames {
		ctx.setFilePermissions(fileName, permissions)
	}
}

func (ctx *AppContext) setFilePermissions(fileName string, permissions os.FileMode) {
	err := os.Chmod(fileName, permissions)
	if err != nil {
		log.Printf("Error setting permissions for %s: %v", fileName, err)
	}
}

func (ctx *AppContext) isProhibited(fileName string) bool {
	for _, name := range ctx.ProtectedNames {
		if strings.Contains(fileName, name) {
			return true
		}
	}
	return false
}

func handleFileOperation(ctx *AppContext) {
	fmt.Println("-------------------------")
	var fileName string
	fmt.Println("Enter operation (create, copy, delete, rename, toggle, exit, add):")
	var operation string
	fmt.Scanln(&operation)

	switch operation {
	case "toggle":
		ctx.toggleProtection()
		fmt.Println("Security Mode Changed")
		handleFileOperation(ctx)
	case "exit":
		fmt.Println("Exiting Program")
		os.Exit(0)
	case "add":
		fmt.Println("Enter the name of the file to add to the prohibited list:")
		var prohibitedName string
		fmt.Scanln(&prohibitedName)
		// Добавляем имя файла в template.tpl
		if err := ctx.addProhibitedNameToFile(prohibitedName); err != nil {
			fmt.Printf("Error adding prohibited name to template.tpl: %v\n", err)
			return
		}
		fmt.Printf("File name '%s' added to the prohibited list.\n", prohibitedName)
		handleFileOperation(ctx)
	default:
		if !isValidOperation(operation) {
			fmt.Printf("Invalid Operation  %s \n", operation)
			return
		}

		fmt.Println("Enter the filename:")
		fmt.Scanln(&fileName)

		if ctx.ProtectionEnabled && ctx.isProhibited(fileName) {
			fmt.Printf("Operation %s on file '%s' is prohibited.\n", operation, fileName)
			return
		}

		switch operation {
		case "create":
			_, err := os.Create(fileName)
			if err != nil {
				fmt.Printf("Error creating file '%s': %v\n", fileName, err)
			} else {
				fmt.Printf("File '%s' created.\n", fileName)
			}
			handleFileOperation(ctx)

		case "copy":
			// Check if the source file exists
			_, err := os.Stat(fileName)
			if err != nil {
				fmt.Printf("Error copying file '%s': Source file does not exist.\n", fileName)
				return
			}

			// Check if the destination file already exists
			_, err = os.Stat("new_" + fileName)
			if err == nil {
				fmt.Printf("Error copying file '%s': Destination file already exists.\n", fileName)
				return
			}

			sourceFile, err := os.Open(fileName)
			if err != nil {
				fmt.Printf("Error opening source file '%s': %v\n", fileName, err)
				return
			}
			defer sourceFile.Close()

			destinationFile, err := os.Create("new_" + fileName)
			if err != nil {
				fmt.Printf("Error creating destination file: %v\n", err)
				return
			}
			defer destinationFile.Close()

			_, err = io.Copy(destinationFile, sourceFile)
			if err != nil {
				fmt.Printf("Error copying file '%s' to 'new_%s': %v\n", fileName, fileName, err)
				return
			}

			fmt.Printf("File '%s' copied to 'new_%s'.\n", fileName, fileName)

		case "delete":
			err := os.Remove(fileName)
			if err != nil {
				fmt.Printf("Error deleting file '%s': %v\n", fileName, err)
				return
			}
			fmt.Printf("File '%s' deleted.\n", fileName)

		case "rename":
			fmt.Println("Enter new file name:")
			var newFileName string
			fmt.Scanln(&newFileName)

			err := os.Rename(fileName, newFileName)
			if err != nil {
				fmt.Printf("Error renaming file '%s' to '%s': %v\n", fileName, newFileName, err)
				return
			}
			fmt.Printf("File '%s' renamed to '%s'.\n", fileName, newFileName)

		default:
			fmt.Printf("Invalid Operation  %s \n", operation)
			return
		}
	}
}

func (ctx *AppContext) addProhibitedNameToFile(fileName string) error {
	// Открываем файл для добавления имени
	file, err := os.OpenFile(ctx.ProtectedFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем имя файла в новую строку
	_, err = file.WriteString(fileName + "\n")
	if err != nil {
		return err
	}

	ctx.ProtectedNames = append(ctx.ProtectedNames, fileName)

	return nil
}

func isValidOperation(operation string) bool {
	// Проверьте, является ли операция допустимой
	validOperations := []string{"create", "copy", "delete", "rename"}
	for _, op := range validOperations {
		if operation == op {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Enter the password to enable/disable protection mode:")
	var userPassword string
	fmt.Scanln(&userPassword)

	appContext := NewAppContext()
	err := appContext.loadProhibitedNames()
	if err != nil {
		log.Printf("Error loading prohibited names: %v", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(appContext.PasswordHash), []byte(userPassword))
	if err != nil {
		fmt.Println("Incorrect password. Protection mode remains unchanged.")
		return
	}

	for {
		handleFileOperation(appContext)
	}
}
