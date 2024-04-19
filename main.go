package main

import (
	"bufio"
	"fmt"
	"os"
)

type Range struct {
	From int
	To   int
}

func AsciiMap(arr []int) map[int]Range {
	res := map[int]Range{}
	for i := range arr {
		if i+1 < len(arr) {
			res[32+i] = Range{From: arr[i+1] - 8, To: arr[i+1] - 1}
		}
		// else if i == len(arr)-1 {
		// 	res[32+1] = Range{From: arr[i-1] - 8, To: arr[i+1] - 1}
		// }
	}
	return res
}

func main() {
	// str := "Boy"
	// Read the commandline arguments
	// arg := os.Args[1:]

	// Read a file
	file, err := os.Open("standard.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	emptyLines := []int{}
	countLine := 1

	// For scanning the file to fill the array where there are empty values
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			emptyLines = append(emptyLines, countLine)
		}
		countLine++

	}

	defer file.Close()

	// Print the arguments
	linesToRead := AsciiMap(emptyLines)
	lineNum := 0

	// scan the file for specific values based on the string and map values
	for scanner.Scan() {
		fmt.Println("called")
		lineNum++
		if rangeToread, ok := linesToRead[lineNum]; ok {
			if lineNum >= rangeToread.From && lineNum <= rangeToread.To {
				fmt.Println(scanner.Text())
			}
		}
	}

	// for i, val := range res {
	// 	fmt.Printf("The ascii values of %d is from : %d to : %d  \n", i, val.From, val.To)
	// }
	// fmt.Println(emptyLines)
	// fmt.Println(res)

	// Handling any scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(linesToRead)
}
