package main

import (
	
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	var str_args string
	if len(args) != 0 {
		str_args += args[0]
	}
	// Open file

	f, err := os.ReadFile("standard.txt")

	if err != nil {
		panic(err)
	}

	

	strArr := strings.Split(string(f), "\n")

	
	var ascii []string
	for i := 0; i < 8; i++ {
		ascii = append(ascii, "")
	}

	for i := 0; i < len(str_args); i++ {
		start_point := (int(str_args[i])-32)*9 + 1
		for i := 0; i < len(ascii); i++ {
			ascii[i] += strArr[start_point+i]
		}

	}
	// fmt.Println(str_args)


	for i := 0; i < len(ascii); i++ {
		fmt.Println(string(ascii[i]))
	}
}
