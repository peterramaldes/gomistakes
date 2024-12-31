package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3} // len=3 cap=3 [1, 2, 3]
	log(s1, "s1")

	s2 := s1[1:2] // len=1 cap=2 [2]
	log(s2, "s2")

	s3 := append(s2, 10) // len=2 cap=2 [2, 10]
	log(s3, "s3")

	log(s1, "s1") // len=3 cap=3 [1, 2, 10]
	log(s2, "s2") // len=1 cap=2 [2]

	// s1: len = 3, cap 3, [1 2 3]
	// s2: len = 1, cap 2, [2]
	// s3: len = 2, cap 2, [2 10]
	// s1: len = 3, cap 3, [1 2 10]
	// s2: len = 1, cap 2, [2]

}

func log(arr []int, env string) {
	fmt.Printf("%s: len = %d, cap %d, %+v\n", env, len(arr), cap(arr), arr)
}
