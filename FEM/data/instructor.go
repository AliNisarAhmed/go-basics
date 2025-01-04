package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

var id int = 0

func increment(id *int) {
	*id++
}

func NewInstructor(name string, lastname string, score int) Instructor {
	increment(&id)
	return Instructor{FirstName: name, LastName: lastname, Id: id, Score: score}
}

func (ins Instructor) Print() string {
	return fmt.Sprintf("%d - %v, %v: (%d)\n", ins.Id, ins.LastName, ins.FirstName, ins.Score)
}
