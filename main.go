package main

import (
	"fmt"
	// "mahi-module/mathOperations"
)


func main() {
	// r := mathOperations.Add(10, 20) //declare and assign the value
	// fmt.Println("Sum : ", r)

	// Example using map
	student := make(map[string]int)
	student["John"] = 90
	student["Alice"] = 85
	student["Bob"] = 92

	fmt.Println("Student:", student)
	fmt.Println("Length of student map:", len(student))

	// Accessing values from the map
	fmt.Println("John's score:", student["John"])
	fmt.Println("Alice's score:", student["Alice"])

	// Updating values in the map
	student["Bob"] = 95
	fmt.Println("Updated Bob's score:", student["Bob"])

	// Deleting a key-value pair from the map
	delete(student, "Alice")
	fmt.Println("Student after deleting Alice:", student)


	
	// To appeand method to the slice
	// First msut be a slice(array) and then the elements to be appended

	// Adding Comments

}


// To run the code, use the 2 ways:
// 1. go run main.go // this will run the code without creating the executable file
// 2. go build main.go // this will create the executable file
