package main

import "fmt"

func main() {
	s := make([]int, 3, 6)
	s[1] = 1
	s = append(s, 2)

	// Note: In Go, a slices grows by doubling size until it contains 1024
	// elements, after that which it grows by 25%
	s = append(s, 3, 4, 5)

	fmt.Printf("s: %v\n", s)

	s1 := make([]int, 3, 6)
	s2 := s1[1:3]

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

	fmt.Println("Updating s1[1] = 1")

	s1[1] = 1
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

	s2 = append(s2, 2)
	fmt.Println("`s2 = append(s2, 2)`")

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

	fmt.Println("Adding elements until the s2 is full")
	s2 = append(s2, 3)
	s2 = append(s2, 4)
	s2 = append(s2, 5)

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))

	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("len(s2): %v\n", len(s2))
	fmt.Printf("cap(s2): %v\n", cap(s2))
}
