package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006"))

	createdDate := time.Date(2020, time.August, 10, 23,23,0,0,time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
