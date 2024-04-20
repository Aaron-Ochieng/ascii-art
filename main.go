package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strings"
)

func main() {

	file, error := os.Open("standard.txt")
	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	

	countLine := 0
	var str string
	for scanner.Scan() {
		line := scanner.Text()
		countLine++
		if countLine >= 1 && countLine <= 856 {
			str += string(line) + "\n"
		}
	
	}

	

strArr := strings.Split(str, "\n")
fmt.Print(len(strArr))

}