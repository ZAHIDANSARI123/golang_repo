package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email   string
	Address string
	Phone   string
}

type Address struct {
	City    string
	Country string
	ZipCode string
}

type Employee struct {
	Person_Details  Person
	Person_Contact  Contact
	Address_Details Address
}

func main() {
	var jahid Person
	fmt.Println("Jahid person :", jahid)
	jahid.FirstName = "Jahid"
	jahid.LastName = "Hossain"
	jahid.Age = 33
	fmt.Println("Jahid person : ", jahid)

	user := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	fmt.Println("User person : ", user)

	var user2 = new(Person)
	user2.FirstName = "Jane"
	user2.LastName = "Smith"
	user2.Age = 28
	fmt.Println("User2 person : ", user2)

	var Employee1 Employee
	Employee1.Person_Details = Person{
		FirstName: "Jahid",
		LastName:  "Hossain",
		Age:       33,
	}
	Employee1.Person_Contact.Email = "jahid@gmail.com"
	Employee1.Person_Contact.Phone = "01700000000"

	Employee1.Address_Details = Address{
		City:    "Dhaka",
		Country: "Bangladesh",
		ZipCode: "1207",
	}

	fmt.Println("Employee1 :", Employee1)
}
