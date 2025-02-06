package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// Normal for loop
	fmt.Println("Normal For Loop")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// For each
	fmt.Println("\nFor each")
	for i := range days {
		fmt.Println(days[i])
	}

	// Same as enumerate sa python
	fmt.Println("\nEnumerate")
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	// Same as enumerate sa python
	fmt.Println("\nEnumerate with blank")
	for _, day := range days {
		fmt.Printf("Value is %v\n", day)
	}


	// While For
	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 5{
			rougeValue++
			continue
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}


}
