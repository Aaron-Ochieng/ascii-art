package main

import (
	"bufio"
	"fmt"
	"os"

	 "strings"
)

func main() {

	file, error := os.Open("standard.txt")
	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	

	countLine := 1
	var str string
	for scanner.Scan() {
		line := scanner.Text()
		countLine++
		if countLine >= 1 && countLine <= 856 {
			str += string(line) + "\n"
		}
	
	}

	

strArr := strings.Split(str, "\n")
//  fmt.Print(len(strArr))

ints := []int{33,34}
str1, str2,str3,str4,str5,str6,str7,str8 := "","","","", "","", "", ""


for i:=0; i < len(ints); i++{
	line_num :=(int(ints[i]) - 32 ) * 9 + 1
	for j:= line_num; j < line_num + 8; j++{
		if j  == line_num {
			str1 += strArr[j]
		}
		if j  == line_num  + 1{
			str2 += strArr[j]
		}
		if j  == line_num + 2 {
			str3 += strArr[j]
		}
		if j  == line_num + 3{
			str4 += strArr[j]
		}
		if j  == line_num + 4{
			str5 += strArr[j]
		}
		if j  == line_num  + 5{
			str6 += strArr[j]
		}
		if j  == line_num + 6 {
			str7+= strArr[j]
		}
		if j  == line_num + 7{
			str8 += strArr[j]
		}
	}
	


}



fmt.Println(str1)
fmt.Println(str2)
fmt.Println(str3)
fmt.Println(str4)
fmt.Println(str5)
fmt.Println(str6)
fmt.Println(str7)
fmt.Println(str8)
}