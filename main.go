package main

import (
	
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	var str_arg string
	if len(args) != 0 {
		str_arg += args[0]
	}
	// Open file
	str_args := strings.Split(str_arg, "\\n")
	f, err := os.ReadFile("standard.txt")

	if err != nil {
		panic(err)
	}

	

	strArr := strings.Split(string(f), "\n")

	


	for i := 0; i < len(str_args); i++ {
		for k := 0; k < 8; k++ {
			for j :=0; j < len(str_args[i]); j++{
				start_point := (int(str_args[i][j])-32)*9 + 1
					fmt.Print( strArr[start_point+k])
			}
			fmt.Println()
		}
	

	}
	// fmt.Println(str_args)


	
}
