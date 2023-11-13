package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if there is at least one command-line argument
	if len(os.Args) < 2{
		fmt.Println("The program should have at least one argument. Please run the program in a correct way.")
		return
	}

	var TextFile string

	// Check if a second command-line argument is provided (for the file name)
 	if len(os.Args) > 2 {
		TextFile = os.Args[2] // Get file name from command-line argument
	} else {
		TextFile = "standard.txt" // Default text file if none is provided
	}

	// Open the text file
	fontFile, err := os.Open(TextFile)
	Handler(err)
	defer fontFile.Close()

	// Get the text to be printed
	text := os.Args[1]

	// Read the text file and store the lines in a slice
	fontList := []string{}
	scanner := bufio.NewScanner(fontFile)
	for scanner.Scan() {
		fontList = append(fontList, scanner.Text())
	}

	// Split the user's input based on the "\\n" delimiter
	splitText := strings.Split(text, "\\n")

	// Iterate through each line in the splitText slice
	for _, line := range splitText {
		// Check if the line is empty
		if len(line) == 0 {
			fmt.Println()
			continue
		}

		// Loop through 8 times (for each row in the characters)
		for i := 0; i < 8; i++ {
			// Loop through each character in the line
			for _, char := range line {
				// Check if the character is within the ASCII range
				if int(char) >= 32 && int(char) <= 126 {
					// Print the corresponding character from fontList
					initial := fontList[i+((int(char)-32)*9)+1]
					fmt.Print(initial)
				}
			}
			fmt.Println() // Move to the next row
		}
	}
}

func Handler(err error) {
	// Check if error is not nil
	if err != nil {
		// Print error message
		fmt.Println(err)
	}
}
