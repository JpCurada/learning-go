package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	jp := User{"JP", "johncurada.02@gmail.com", true, 20}
	fmt.Println(jp)
	fmt.Printf("JP details are: %+v\n", jp)
}
