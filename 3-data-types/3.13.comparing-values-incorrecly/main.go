package main

import (
	"fmt"
)

type customer struct {
	id         string
	operations []float64
}

func (a customer) equal(b customer) bool {
	if a.id != b.id {
		return false
	}

	if len(a.operations) != len(b.operations) {
		return false
	}

	for i := 0; i < len(a.operations); i++ {
		if a.operations[i] != b.operations[i] {
			return false
		}
	}

	return true
}

func main() {
	c1 := customer{id: "x", operations: []float64{1.}}
	c2 := customer{id: "x", operations: []float64{2.}}
	fmt.Println(c1.equal(c2))
}
