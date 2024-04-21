package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var str_args string
	if len(args) != 0 {
		str_args += args[0]
	}
	// Open file

	f, err := os.Open("standard.txt")

	if err != nil {
		panic(err)
	}

	scanning := bufio.NewScanner(f)

	var strArr []string

	for scanning.Scan() {
		strArr = append(strArr, scanning.Text())
	}
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
	fmt.Println(strArr)
	fmt.Println("Len>>>>", len(strArr))
	fmt.Println(">>>", ascii)

	for i := 0; i < len(ascii); i++ {
		fmt.Println(string(ascii[i]))
	}
}
