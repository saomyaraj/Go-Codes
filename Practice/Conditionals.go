package main

import "fmt"

func main() {
	height := 10
	threshold := 5
	if height > threshold {
		fmt.Println("You are tall enough")
	} else if height < threshold {
		fmt.Println("You are short enough")
	} else {
		fmt.Println("You are at perfect size")
	}

	//Other way of conditionals
	//email := "xyz123@gmail.com"
	//if length := getLength(email); length<1 {
	//	fmt.Println("Email is invalid")
	//}
}
