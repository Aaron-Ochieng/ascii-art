// Ascii-art is a program which consists in receiving a string as an argument
// and outputting the string in a graphic representation using ASCII.
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	ascii "ascii/utilities"
	"ascii/colors"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <input string>")
		return
	}

	var input string
	if len(os.Args[1:]) > 2 {
		input = os.Args[4] // user input

	} else {
		input = os.Args[2]
	}
	// fmt.Println(len(input))

	var color string
	var letterToBeColored string
	letters := false

	flag.StringVar(&color, "color", "", "the color desired by the user")
	flag.StringVar(&letterToBeColored, "l", "", "the letter or letters that you can chose to be colored.")

	flag.Parse()

	pattern := `^--color=.+(.+)?(.)$`

	re := regexp.MustCompile(pattern)

	arguments := strings.Join(os.Args[1:], " ")

	if !re.MatchString(arguments) {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=< CAN'T BE EMPTY > <letters to be colored> something")
		os.Exit(0)
	} 

	if color == ""{
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=< CAN'T BE EMPTY > <letters to be colored> something")
		os.Exit(0)
	}

	if letterToBeColored != "" {
		letters = true
	}

	if input == "" {
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

	count := 0
	ansiColor := utils.GetColor(color)

	for _, part := range inputParts {
		if part == "" {
			count++
			if count < len(inputParts) {
				fmt.Println() // Print a newline if the part is empty (i.e., consecutive newline characters)
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ { // this loop is responsible for the height of each character
			for _, char := range part { // iterates through each character of the current word
				startingIndex := ascii.GetStartingIndex(int(char)) // obtaining the starting position of the char
				line := fileData[startingIndex+i]
				if startingIndex >= 0 {
					if letters {
						if strings.ContainsRune(letterToBeColored, char) {
							fmt.Printf("%s%s\x1b[0m", ansiColor, line)
						} else {
							fmt.Print(line)
						}
					}else {

						fmt.Printf("%s%s\x1b[0m", ansiColor, line)// printing the character line by line
					}
				}
			}
			fmt.Println() // printing a new line after printing each line of the charcter
		}
	}
}
