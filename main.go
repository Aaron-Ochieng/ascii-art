package main

import (
	"bufio"
	"fmt"
	"os"
	
)

func main() {
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	const startLine, endLine = 1, 856
	var selectedLines []string
	for i := startLine - 1; i < endLine && i < len(lines); i++ {
		selectedLines = append(selectedLines, lines[i])
	}

	ints := []int{33, 34, 35}
	var strs []string
	for _, num := range ints {
		lineNum := (num - 32) * 9
		for j := 0; j < 8 && lineNum+j < len(selectedLines); j++ {
			if len(strs) <= j {
				strs = append(strs, "")
			}
			strs[j] += selectedLines[lineNum+j]
		}
	}

	for _, str := range strs {
		fmt.Println(str)
	}
}
