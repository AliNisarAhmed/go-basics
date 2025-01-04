package data

import "fmt"

type Duration float32 // in hours

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (w Course) SignUp() bool {
	return true
}

func (c Course) String() string {
	return fmt.Sprintf("--- %v --- (by %v)", c.Name, c.Instructor.FirstName)
}
