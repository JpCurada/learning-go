package main

import "fmt"

func main() {
	// SECTION 1: String Variables
	var strFirstName string = "John Paul"  // Explicit type
	var strLastName = "Curada"            // Type inference
	var strName string                    // Zero value initialization
	
	fmt.Println("Initial names:", strFirstName, strLastName, strName)
	
	strName = strLastName
	strLastName = "Mercado"
	fmt.Println("After modification:", strFirstName, strLastName, strName)

	// SECTION 2: Integer Variables
	var intAgeOne int = 20    // Explicit type
	var intAgeTwo = 30        // Type inference
	intAgeThree := 40         // Short declaration
	fmt.Println("\nIntegers:", intAgeOne, intAgeTwo, intAgeThree)

	// Integer types with specific sizes
	var numOne int8 = 25              // 8-bit signed integer (-128 to 127)
	var numTwo int8 = -123
	var absOne uint8 = 12             // 8-bit unsigned integer (0 to 255)
	var absTwo uint32 = 1213112431    // 32-bit unsigned integer
	fmt.Println("Different integer types:", numOne, numTwo, absOne, absTwo)

	// SECTION 3: Floating Point Variables
	var scoreOne float64 = 10.21               // Explicit float64
	var scoreTwo = -12.03213123213213213       // Type inference
	scoreThree := 12.213123213                 // Short declaration
	fmt.Println("\nFloating points:", scoreOne, scoreTwo, scoreThree)

	// SECTION 4: Print Function Variations
	fmt.Print("Hello, ")           // Print without newline
	fmt.Print("World! \n")        // Manual newline
	fmt.Print("New line example \n")

	fmt.Println("\nPrintln examples:")  // Println automatically adds newline
	fmt.Println("Hello ninjas!")
	fmt.Println("goodbye ninjas!")

	// SECTION 5: Formatted Printing (Printf)
	name := "JP"
	age := 20

	// Basic variable substitution
	fmt.Println("\nFormatted print examples:")
	fmt.Printf("Hello %v! Are you already %v years old?\n", name, age)
	
	// String formatting with quotes
	fmt.Printf("Hello %q! Are you already %q years old?\n", name, age)
	
	// Type printing
	fmt.Printf("Variable type: %T\n", age)
	
	// Float formatting
	fmt.Printf("Float with 1 decimal: %0.1f\n", 225.27)

	// SECTION 6: Sprintf (String Formatting)
	savedString := fmt.Sprintf("Hello %v! Are you already %v years old?\n", name, age)
	fmt.Println("\nStored formatted string:", savedString)
}