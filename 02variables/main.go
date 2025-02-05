package main

import "fmt"

// Make sure to capitalize the first letter of the variable
const LoginToken string = "mamamo123"

func main() {

	var strUsername string = "jpcurada_"
	fmt.Println(strUsername)
	fmt.Printf("Variable is of type: %T \n", strUsername)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)


	var intVal uint8 = 255
	fmt.Println(intVal)
	fmt.Printf("Variable is of type: %T \n", intVal)

	var floatVal float64 = 255.231232312
	fmt.Println(floatVal)
	fmt.Printf("Variable is of type: %T \n", floatVal)

	var anotherVariable float32
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var website = "learnoeonline.in"
	fmt.Println(website)

	// no var style (walrus operator is not allowed outside mthods)
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
