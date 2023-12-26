package magazine

import "fmt"

type Subcsriber struct {
	Name   string
	Rate   float64
	Active bool
}

func DefaultSubcriber(name string) *Subcsriber {
	var sub Subcsriber
	sub.Name = name
	sub.Rate = 5.99
	sub.Active = true
	return &sub
}
func PrintInfo(sub *Subcsriber) {
	fmt.Println("Name:", sub.Name)
	fmt.Println("Mounthly rate:", sub.Rate)
	fmt.Println("Active:", sub.Active)
}

func ApplyDiscount(s *Subcsriber) {
	s.Rate = 4.99
}

type Employee struct {
	Name   string
	Salary float64
}
