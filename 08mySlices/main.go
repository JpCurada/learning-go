package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "tomato", "peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fmt.Println(fruitList[1:3])

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// Allocate and override the limit 4
	highScores = append(highScores, 555, 666, 321)

	// Sort
	fmt.Println(sort.IntsAreSorted(highScores)) // return false
	sort.Ints(highScores)

	fmt.Println(highScores)

	// How to remove a value from slices based on index
	var arrNames = []string{"Jp", "Pj", "John", "PJ"}
	var index int = 2;
	arrNames = append(arrNames[:index], arrNames[index+1:]...)
	fmt.Println(arrNames)


}
