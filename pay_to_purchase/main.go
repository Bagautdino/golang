package main

import (
	"bufio"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"strings"
	"time"
)

const (
	registryKey    = `SOFTWARE\filler`
	dataFile       = "user_data.txt"
	maxLaunches    = 5
	maxDuration    = 3 * time.Minute
	launchLimitMsg = "You have reached the launch limit. Please purchase the full version or uninstall the program."
	timeLimitMsg   = "You have reached the time limit. Please purchase the full version or uninstall the program."
)

func main() {
	previousUsage := readRegistry()
	fmt.Println("Welcome to the User Data Program!")

	// Check if the user data file exists
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		fmt.Println("No user data found.")
	}

	launchCount := 0
	startTime := time.Now()

	for {
		elapsedTime := time.Since(startTime)
		if elapsedTime >= maxDuration {
			fmt.Println(timeLimitMsg)
			promptPurchaseOrUninstall()
			break
		}

		if launchCount >= maxLaunches {
			fmt.Println(launchLimitMsg)
			promptPurchaseOrUninstall()
			break
		}

		fmt.Print("Enter your last name (or 'q' to quit): ")
		lastName := getUserInput()

		if strings.ToLower(lastName) == "q" {
			break
		}

		fmt.Print("Enter your first name: ")
		firstName := getUserInput()

		fmt.Print("Enter your middle name: ")
		middleName := getUserInput()

		fullName := fmt.Sprintf("%s %s %s", lastName, firstName, middleName)
		saveRegistry(previousUsage)
		if !isLaunchLimitReached(launchCount) {
			if isFullNameInFile(fullName) {
				fmt.Println("The entered full name exists in the data file.")
			} else {
				saveUserData(fullName)
			}

			launchCount++
			fmt.Printf("Saved: %s\n", fullName)
		} else {
			fmt.Println(launchLimitMsg)
			promptPurchaseOrUninstall()
			break
		}
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func saveUserData(fullName string) {
	file, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening data file:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, fullName)
	if err != nil {
		fmt.Println("Error writing to data file:", err)
	}
}

func isLaunchLimitReached(launchCount int) bool {
	return launchCount >= maxLaunches
}

func isFullNameInFile(fullName string) bool {
	file, err := os.Open(dataFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == fullName {
			return true
		}
	}

	return false
}

func promptPurchaseOrUninstall() {
	fmt.Println("Options:")
	fmt.Println("1. Purchase the full version")
	fmt.Println("2. Uninstall the program")
	fmt.Print("Enter your choice (1 or 2): ")
	choice := getUserInput()

	switch choice {
	case "1":
		fmt.Println("Please visit the following link to access the full version:")
		fmt.Println("https://example.com/full-version-download") // Replace with the actual download link
		// Add code here to provide full version features.
	case "2":
		fmt.Println("Uninstalling the program...")

		err := os.Remove(os.Args[0]) // os.Args[0] contains the path to the executable
		if err != nil {
			fmt.Println("Error uninstalling the program:", err)
			return
		}

		fmt.Println("Program uninstalled successfully.")
		os.Exit(0) // Exit the program
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}

func readRegistry() int {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, registryKey, registry.READ)
	if err != nil {
		return 0
	}
	defer key.Close()

	usage, _, err := key.GetIntegerValue("Usage")
	if err != nil {
		return 0
	}

	return int(usage)
}

func saveRegistry(usage int) {
	key, _, err := registry.CreateKey(registry.LOCAL_MACHINE, registryKey, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Error creating or opening the registry key:", err)
		return
	}
	defer key.Close()

	err = key.SetDWordValue("Usage", uint32(usage))
	if err != nil {
		fmt.Println("Error setting registry value:", err)
	}
}
