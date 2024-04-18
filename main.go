package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Read the commandline arguments
	// arg := os.Args[1:]

	// Read a file
	file, error := os.Open("standard.txt")
	if error != nil {
		fmt.Println(error)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	emptyLines := []int{}
	countLine := 1
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			emptyLines = append(emptyLines, countLine)
		}
		countLine++

	}

	defer file.Close()

	fmt.Println(emptyLines)
	// Print the arguments
	res := AsciiMap(emptyLines)
	fmt.Println(res)

}
