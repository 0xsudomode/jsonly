package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func printUsage() {
	fmt.Println("Usage: jsonly  < input_file")
	fmt.Println("Filters URLs that end with .js from stdin input.")
}

func main() {
	// Check if there are command-line arguments
	if len(os.Args) > 1 {
		printUsage()
		os.Exit(1)
	}

	// Regular expression for matching URLs
	urlPattern := `^(https?://[^\s]+)$`
	jsExtensionPattern := `\.js$`

	// Compile the regular expressions
	urlRegex := regexp.MustCompile(urlPattern)
	jsExtensionRegex := regexp.MustCompile(jsExtensionPattern)

	// Check if there is input from stdin
	if !isStdinAvailable() {
		printUsage()
		os.Exit(1)
	}

	// Read input from stdin line by line
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line matches the URL pattern and ends with ".js"
		if urlRegex.MatchString(line) && jsExtensionRegex.MatchString(line) {
			fmt.Println(line)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

// Function to check if stdin has data
func isStdinAvailable() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
		return false
	}

	// Check if stdin has data (is not empty)
	return (info.Mode()&os.ModeCharDevice) == 0
}
