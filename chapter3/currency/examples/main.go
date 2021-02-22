package main

import (
	"fmt"
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter3/currency"
)

func main() {
	userInput := "15.93"

	pennies, err := currency.CovertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User input converted to %d penies \n", pennies)

	pennies += 15

	dollars := currency.ConvertPenniesToDollerString(pennies)

	fmt.Printf("Add 15 cents, new values is %s dollars\n", dollars)

}
