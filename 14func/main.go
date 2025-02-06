package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	header()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your height in meters: ")
	strHeightInput, _ := reader.ReadString('\n')

	fmt.Println("Please enter your weight in kilograms: ")
	strWeightInput, _ := reader.ReadString('\n')

	floatHeight, floatWeight := cleanInput(strHeightInput, strWeightInput)
	floatBMI := computeBMI(floatHeight, floatWeight)

	fmt.Printf("You're BMI is %.2f\n\n", floatBMI)
	classifyBMI(floatBMI)

	averageBMI := getAverageBMI(floatBMI, floatBMI, floatBMI)
	fmt.Printf("Average BMI: %2f", averageBMI)
}

// Normal functions (no parameters, no return)
func header() {
	fmt.Println("Welcome to BMI Calculator ")
}

// Function with parameters but no return output
func classifyBMI(floatBMI float64) {
	if floatBMI <= 18 {
		fmt.Println("BMI Classification: Underweight")
	} else if floatBMI < 25 {
		fmt.Println("BMI Classification: Normal")
	} else if floatBMI < 30 {
		fmt.Println("BMI Classification: Overweight")
	} else {
		fmt.Println("BMI Classification: Obese")
	}
}

// Function with paramters and one return output
func computeBMI(floatHeight float64, floatWeight float64) float64 {
	floatBMI := floatWeight / (floatHeight * floatHeight)
	return floatBMI
}

// Function with two outputs
func cleanInput(strHeightInput string, strWeightInput string) (float64, float64) {
	floatHeight, _ := strconv.ParseFloat(strings.TrimSpace(strHeightInput), 64)
	floatWeight, _ := strconv.ParseFloat(strings.TrimSpace(strWeightInput), 64)
	return floatHeight, floatWeight
}

// Function for slice
func getAverageBMI(sliceBMI ...float64) float64{
	floatSumBMIs := 0.00
	intlenBMISlice := len(sliceBMI)
	for i := range sliceBMI {
		floatSumBMIs += sliceBMI[i]
	}

	AverageBMI := floatSumBMIs / float64 (intlenBMISlice)
	return AverageBMI
}
