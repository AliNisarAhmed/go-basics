package main

import (
	"ali.com/go/fem/data"
	"fmt"
)

func main() {

	ali := data.NewInstructor("Ali", "Ahmed", 100)
	azlan := data.NewInstructor("Azlan", "Ali", 200)

	goCourse := data.Course{Id: 1, Name: "Go Fundamentals", Instructor: ali}

	swiftWS := data.NewWorkshop("Swift Workshop", azlan)

	print(ali.Print())
	print(azlan.Print())
	fmt.Printf("%v\n", goCourse)
	fmt.Printf("%v\n", swiftWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}
}
