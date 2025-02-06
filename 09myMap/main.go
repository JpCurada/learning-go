package main

import "fmt"

func main() {
	fmt.Println("Maps of Array (Dictionary of List)")
	var courseMap = make(map[string][]string)
	courseMap["CCIS"] = []string{"BSCS", "BSIT"}
	courseMap["CEA"] = []string{"BSCE", "BSME", "BSECE", "BSCmpE", "BSEE", "BSIE"}
	
	fmt.Println(len(courseMap["CCIS"]))
	fmt.Println(len(courseMap["CEA"]))
}


