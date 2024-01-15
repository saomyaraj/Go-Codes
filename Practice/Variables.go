package main

import "fmt"

func main() {
	//Initialize the variables
	var smsSendLimit int
	var costPerSMS float64
	var hasPermission bool
	var userName string
	fmt.Printf(
		"%v %f %v %q\n",
		smsSendLimit,
		costPerSMS,
		hasPermission,
		userName,
	)

	//Assigning the value to the variables
	greet := "Hello Saomyaraj"
	numCars := 10
	temperature := 12.6
	isFunny := true
	fmt.Println(greet)
	fmt.Println("The numbers of cars is ", numCars)
	fmt.Println(temperature)
	fmt.Println(isFunny)

	averageOpenRate, displayMessage := .23, "is the average open rate of your message"
	fmt.Println(averageOpenRate, displayMessage)

	tempInt := int(temperature)
	fmt.Println(tempInt)

	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

}
