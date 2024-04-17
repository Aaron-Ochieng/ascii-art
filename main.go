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
	readfile, error := os.Open("standard.txt")
	if error != nil {
		fmt.Println(error)
	}

	fileScanner := bufio.NewScanner(readfile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text() + fileScanner.Text())

	}

	defer readfile.Close()

}
