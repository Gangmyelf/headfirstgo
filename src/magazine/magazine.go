package magazine

import "fmt"

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Subcsriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
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
