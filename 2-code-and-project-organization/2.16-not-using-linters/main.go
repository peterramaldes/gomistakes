package main

import (
	"fmt"
)

func main() {
	i := 1
	if true {
		i := 2 // Shadowing (linter should cover this)
		fmt.Println(i)
	}
	fmt.Println(i)
}
