package main

import "fmt"

func init() {
	fmt.Println("Hello init 1")
}

func init() {
	fmt.Println("Hello init 2")
}

func main() {
	stateTax, cityTax := fmt.Println(calculate(100))
	fmt.Println(cityTax)
	fmt.Println(stateTax)
}

func calculate(price float32) (float32, float32) {
	return price * 0.09, price * 0.1
}

func calculateWithName(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return
}
