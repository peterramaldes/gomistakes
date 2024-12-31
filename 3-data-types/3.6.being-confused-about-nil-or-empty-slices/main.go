package main

import "fmt"

func main() {
	var s []string
	log(1, s)

	s = []string(nil)
	log(2, s)

	s = []string{}
	log(3, s)

	s = make([]string, 0)
	log(4, s)

	s = f()
	log(5, s)
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}

func f() []string {
	var s []string
	if foo() {
		s = append(s, "foo")
	}
	if bar() {
		s = append(s, "bar")
	}
	return s
}

func bar() bool { return false }
func foo() bool { return false }
