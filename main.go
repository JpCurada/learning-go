package main

import "fmt"

func main() {
    // SECTION 1: ARRAYS (Fixed Length)
    fmt.Println("=== Arrays ===")
    
    // Array declaration with explicit size
    var arrNums = [3]int{20, 25, 30}
    fmt.Printf("Integer array: %v, Length: %d\n", arrNums, len(arrNums))

    // Array with implicit size
    names := [...]string{"yoshi", "mario", "peach", "bowser"}
    fmt.Printf("Original names: %v, Length: %d\n", names, len(names))

    // Modifying array elements
    names[1] = "JP"
    fmt.Printf("Modified names: %v\n", names)

    // Zero-valued array
    var emptyArr [3]int
    fmt.Printf("Zero-valued array: %v\n\n", emptyArr)

    // SECTION 2: SLICES (Dynamic Length)
    fmt.Println("=== Slices ===")
    
    // Basic slice declaration
    var scores = []int{100, 50, 60}
    fmt.Printf("Original slice: %v, Length: %d\n", scores, len(scores))

    // Modifying and appending to slices
    scores[2] = 25
    scores = append(scores, 85)
    fmt.Printf("Modified slice: %v, Length: %d\n", scores, len(scores))

    // SECTION 3: SLICE OPERATIONS
    fmt.Println("\n=== Slice Operations ===")
    
    fruits := [6]string{"apple", "banana", "orange", "grape", "mango", "kiwi"}
    fmt.Printf("Original array: %v\n", fruits)

    // Different ways of slicing
    rangeOne := fruits[1:3]    // Elements from index 1 to 2
    rangeTwo := fruits[2:]     // Elements from index 2 to end
    rangeThree := fruits[:3]   // Elements from start to index 2
    rangeFour := fruits[:]     // All elements
    
    fmt.Printf("Slice [1:3]: %v\n", rangeOne)
    fmt.Printf("Slice [2:]: %v\n", rangeTwo)
    fmt.Printf("Slice [:3]: %v\n", rangeThree)
    fmt.Printf("Slice [:]: %v\n", rangeFour)
}