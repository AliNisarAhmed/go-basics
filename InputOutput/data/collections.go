package data

import "fmt"

var Countries [3]string

var Slice []int
var Codes map[int]string

// there can be more than 1 init functions, even in same package
func init() {
	Countries[0] = "Canada"
	Countries[1] = "US"
	Countries[2] = "Israel"
	
  qty := len(Countries)

  wellknownPorts := map[string]int { "http": 80, "https": 443}
  fmt.Println(qty)
  fmt.Println(wellknownPorts)
}

func Print_countries() {
  fmt.Println(Countries[0])
}
