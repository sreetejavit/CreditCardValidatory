package main

//https://go.dev/doc/tutorial/call-module-code

import (
	"fmt"

	c "example.com/cred/creditvalidty"
)

func main() {
	fmt.Println("Welcome to Credit Card Validator.")
	cardNumberGiven := &[]int{7, 9, 9, 2, 7, 3, 9, 8, 7, 9, 3}
	err := c.CheckCard(cardNumberGiven)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(" Card is Valid")
	}

}
