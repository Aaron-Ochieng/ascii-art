package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func errHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// generating the starting number on the ascii art file
func getStartingIndex(ascii int) int {
	return (ascii-32)*9 + 2
}

func handleBackspace(s string) string {
	index := strings.Index(s, "\\b") // obtaining the index of the first occurance of \b character

	if index == -1 { // Base case: no "\b" found, return the string as is
		return s
	}
	if index == 0 { // If "\b" is at the beginning, remove it and return the rest of the string
		return handleBackspace(s[1:])
	}
	return handleBackspace(s[:index-1] + s[index+2:]) // Remove the "\b" and the character before it recursively
}

func checkIllegalChar(arr []string) ([]string, error) {
	for _, s := range arr {
		for _, r := range s {
			if r < 32 || r > 126 {
				return []string{}, errors.New("the string must contain characters in the range 32 - 126")
			}
		}
	}
	return arr, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input string>")
		return
	}

	input := os.Args[1] // user input
	fmt.Println(len(input))

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}
	input = handleBackspace(input)
	input = strings.ReplaceAll(string(input), ">", "\\n")   // handling the tab sdoks sceequence
	input = strings.ReplaceAll(string(input), "\\t", "   ") // handling the tab sequence
	file, err := os.ReadFile("standard.txt")
	errHandler(err)

	fileData := strings.Split(string(file), "\n")

	spString := strings.Split(input, "\\n") // Split input by "\\n" to handle newline sequence
	// check for  illegal characters in the string
	inputParts, err := checkIllegalChar(spString)
	errHandler(err)

	for _, part := range inputParts {
		if part == "" {
			fmt.Println() // Print a newline if the part is empty (i.e., consecutive newline characters)
			continue
		}
		for i := 0; i < 8; i++ { // this loop is responsible for the height of each character
			for _, char := range part { // iterates through each character of the current word
				startingIndex := getStartingIndex(int(char)) // obtaining the starting position of the char
				fmt.Print(fileData[startingIndex-1+i])       // printing the character line by line
			}
			fmt.Println() // printing a new line after printing each line of the charcter
		}
	}
}
