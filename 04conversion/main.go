package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our BMI Calculator")

	reader := bufio.NewReader(os.Stdin)

	// Weight: Keep reading until `\n`
	fmt.Println("Please enter your weight in kilograms")
	floatWeightInput, _ := reader.ReadString('\n')
	floatWeight, err := strconv.ParseFloat(strings.TrimSpace(floatWeightInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("You've entered %.2fkg \n", floatWeight)
	}

	// Height: Keep reading until `\n`
	fmt.Println("Please enter your Height in meters")
	floatHeightInput, _ := reader.ReadString('\n')
	floatHeight, err := strconv.ParseFloat(strings.TrimSpace(floatHeightInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("You've entered %.2fm \n", floatHeight)
	}

	floatBMI := floatWeight / (floatHeight * floatHeight)
	fmt.Printf("BMI: %.2f \n", floatBMI)
	

}
