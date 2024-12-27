package main

import (
	"fmt"
	"strconv"
)

type customConstraint interface {
	~int
	String() string
}

func getKeys[K comparable, V any](m map[K]V) []K {
	return []K{}
}

type customInt int

func (i customInt) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	m := map[customInt]string{1: "bar"}
	fmt.Printf("getKeys(m): %v\n", getKeys(m))
}
