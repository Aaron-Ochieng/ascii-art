package main

import (
	"fmt"
	"os"
	"strings"
)

func errHandler(err error) {
	if err != nil {
		fmt.Println("Error", err)
	}
}

func stringContain(s string) []string {
	if strings.Contains(s, "o") {
		return strings.Split(s, "\r\n")
	}
	return strings.Split(s, "\n")
}

func GetIndex(ascii int) int {
	return (ascii-32)*9 + 1
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <inputstring>")
	}
	input := os.Args[1]
	// fmt.Println(input)

	if input == "" {
		return
	}

	file, err := os.ReadFile("standard.txt")
	errHandler(err)
	// fmt.Println(len(file))

	fileData := stringContain(string(file))
	inputData := strings.Split(input, "\\n")

	for _, word := range inputData {
		for i := 0; i < 8; i++ {
			for _, char := range word {
				// fmt.Print(string(char))
				startingIndex := GetIndex(int(char))
				fmt.Print(fileData[startingIndex+i])
			}
			fmt.Println()
		}
	}

	// fmt.Println(len(fileData))
}
