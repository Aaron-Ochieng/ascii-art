package main

import (
	"bufio"
	"fmt"
	"os"

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

	

//  strArr := strings.Split(str, "\n")
//  fmt.Print(len(strArr))

ints := []int{45,46}


var intArr []int
for i:=0; i < len(ints); i++{
	line_num :=(int(ints[i]) - 32 ) * 9 + 1
	
intArr = append(intArr, line_num)

}


fmt.Println(intArr)
}