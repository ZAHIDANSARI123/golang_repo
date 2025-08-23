package main


import "fmt"


func main() {
	x := 22
	if x > 5 {
		fmt.Println("x is greter then 5")
	} else {
		fmt.Println("x is smaller then 5")
	}


	


	z := 10
	if z > 10 {
		fmt.Println("z is greater then 10")
	} else if z > 5 {
		fmt.Println("z is greater then	5 but smaller then 10")
	} else {
		fmt.Println("z is smaller then 5")
	}


	y := 10
	if x < 5 && (y > 5 || z < 44) {
		fmt.Println("hey how are you ? ")
	} else {
		fmt.Println("Learn programming with hello world")
	}
}