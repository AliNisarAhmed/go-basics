package data

import "time"

type Workshop struct {
	Course // Course is embedded within Workshop
	Instructor
	Name string
	Date   time.Time
}

func (w Workshop) SignUp() bool {
  return true
}

func NewWorkshop(name string, instr Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	// w.Course.Name = "Course Name"
	w.Instructor = instr
	return w
}
