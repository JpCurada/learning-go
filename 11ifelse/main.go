package main

import "fmt"

func main() {
	fmt.Println("IF-ELSE Statements")

	/* 

	SYNTAX

	if <condition> {
		action
	} else if <condition> {
	 	action
	} else {
		action 
	}
	*/

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Something Else"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 10%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	// Sometimes value is going out from the web request
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
}
