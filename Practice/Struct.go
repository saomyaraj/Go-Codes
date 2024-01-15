package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("-----------------")
}

func main() {
	test(messageToSend{
		phoneNumber: 8888888888,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 83972438275,
		message:     "Love to see you aboard",
	})
	test(messageToSend{
		phoneNumber: 71346617974,
		message:     "We are excited",
	})
}
