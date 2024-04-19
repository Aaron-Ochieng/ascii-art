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
	countLine := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			emptyLines = append(emptyLines, countLine)
		}
		countLine++

	}

	defer file.Close()

	fi, error := os.Open("standard.txt")
	if error != nil {
		fmt.Println(error)
	}
	scans := bufio.NewScanner(fi)

	// Print the arguments
	res := AsciiMap(emptyLines)
	ints := []int{33, 34, 35}

	coun := 0

	for scans.Scan() {
		line := scans.Text()
		coun++

		for i := 0; i < len(ints); i++ {
			x := res[ints[i]]

			if coun >= x.From && coun <= x.To {
				fmt.Println(line)
			}
		}

	}

	// for i, val := range res {
	// 	fmt.Printf("The ascii values of %d is from : %d to : %d  \n", i, val.From, val.To)
	// }
}
