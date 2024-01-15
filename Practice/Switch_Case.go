package main

import "fmt"

func printNumericalValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Println("%T\n", v)
	case string:
		fmt.Println("%T\n", v)
	default:
		fmt.Println("%T\n", v)
	}
}

func main() {
	printNumericalValue(1)
	printNumericalValue("1")
	printNumericalValue(struct{}{})
}
