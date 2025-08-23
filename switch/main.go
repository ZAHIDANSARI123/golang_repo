package main


import "fmt"


func main() {
	day := 3


	switch day {
	case 1 : 
		fmt.Println("Monday")
		// fmt.Println("Good Day")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown Day")
	}



	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("winter")

	case "April", "May", "June":
		fmt.Println("Spring")
	default: 
		fmt.Println("Other season")
	}


	temperature := 12
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default: 
		fmt.Println("Hot")
	}
}
