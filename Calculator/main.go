package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR")
	fmt.Println("========================")
	fmt.Println("Enter the operator: (+, -, /, *)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number: ")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter second number: ")
	fmt.Scanf("%d", &number2)

	switch operation {
	case "+":
		fmt.Printf("%d + %d = %d\n", number1, number2, number1+number2)
	case "-":
		fmt.Printf("%d - %d = %d\n", number1, number2, number1-number2)
	case "*":
		fmt.Printf("%d * %d = %d\n", number1, number2, number1*number2)
	case "/":
		fmt.Printf("%d / %d = %0.2f\n", number1, number2, float64(number1/number2))
	default:
		panic("Invalid operation")
	}
}
