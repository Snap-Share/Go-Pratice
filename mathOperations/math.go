package mathOperations

import (
	"fmt"
	os "os"
)

func Add(num1, num2 int) int {
	return num1 + num2
}

func MathOperations() {
	// Example of math operations
	var num int = 10 // Declaring a variable with showing the type
	fmt.Println("Number:", num)
	num1 := 10 		// Declaring a variable without showing the type
	num2 := 5

	// Addition
	sum := num1 + num2
	fmt.Println("Sum:", sum)
	fmt.Printf("Sum: %v", sum) // Another way to print the output
	// Convert sum to string and print it using fmt.Fprintf
	sumStr := fmt.Sprintf("%d", sum)
	fmt.Fprintf(os.Stdout, "Sum as string: %s\n", sumStr)
	// Subtraction
	diff := num1 - num2
	fmt.Println("Difference:", diff)

	// Multiplication
	product := num1 * num2
	fmt.Println("Product:", product)

	// Division
	quotient := num1 / num2
	fmt.Println("Quotient:", quotient)

	// Modulo
	remainder := num1 % num2
	fmt.Println("Remainder:", remainder)
}


// Varialbe and Function names are exported 
//if they start with a capital letter