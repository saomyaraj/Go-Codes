package main

import "fmt"

func sub(x int, y int) int {
	return x - y
}

func add(x, y int) int {
	return x + y
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

func getNames() (string, string) {
	return "Saomyaraj", "Jha"
}

func main() {
	diff := sub(4, 3)
	sum := add(3, 4)
	fmt.Println("The difference of two numbers is ", diff)
	fmt.Println("The addition of two numbers is ", sum)

	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")

	firstName, lastName := getNames() //You can replace lastname with _ if you don't want to use it later.
	fmt.Println("Welcome to Tokyo,", firstName, lastName)
}
