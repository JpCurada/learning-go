package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	strWelcome := "Welcome to user input"
	fmt.Println(strWelcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our course:")

	// comma ok || err ok
	// ok, error
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating: %T ", input)

}
