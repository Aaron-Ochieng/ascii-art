// Ascii-art is a program which consists in receiving a string as an argument
// and outputting the string in a graphic representation using ASCII.
package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/utilities"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input string>")
		return
	}

	input := os.Args[1] // user input
	// fmt.Println(len(input))

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}
	input = ascii.HandleBackspace(input)
	input = strings.ReplaceAll(string(input), "\\t", "   ") // handling the tab sequence

	// Read the ascii art text file
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	fileData := ascii.StringContain(string(file))

	// Handling empty files
	if len(fileData) != 856 {
		fmt.Println("Error: >> Banner files  is empty with length of: ", len(file))
		return
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	inputParts, err := ascii.HandleNewLine(input)
	ascii.ErrHandler(err)

	for _, part := range inputParts {
		if part == "" {
			fmt.Println() // Print a newline if the part is empty (i.e., consecutive newline characters)
			continue
		}
		for i := 0; i < 8; i++ { // this loop is responsible for the height of each character
			for _, char := range part { // iterates through each character of the current word
				startingIndex := ascii.GetStartingIndex(int(char)) // obtaining the starting position of the char
				if startingIndex >= 0 {
					fmt.Print(fileData[startingIndex+i]) // printing the character line by line
				}
			}
			fmt.Println() // printing a new line after printing each line of the charcter
		}
	}
}
