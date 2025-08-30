package main

import ("fmt"
	"time"
)


func main() {

	currentTime := time.Now()
	fmt.Println("Current Time: ", currentTime)
	fmt.Printf("Type of currentTime: %T\n", currentTime)

	// currentTime fetch
	formattedTime := currentTime.Format("02-01-2006, Monday, 15:04:05, 03:04 PM")
	fmt.Println("Formatted Time: ", formattedTime)


	layout_str := "02-01-2007"
	dataStr := "25-11-2030"
	formatted_time, _ := time.Parse(layout_str, dataStr)
	fmt.Println("Formatted Time from String: ", formatted_time)
}
