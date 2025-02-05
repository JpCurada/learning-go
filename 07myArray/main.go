package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golans")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Banana"
	// fruitList[3] = "Peach"
	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Length :", len(fruitList))

	var arrnames = [4]string{"Jp", "john", "Paul", "mai"}
	fmt.Println("Length :", len(arrnames))


}
