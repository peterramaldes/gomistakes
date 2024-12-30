package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int32 = math.MaxInt32
	// This operation doesn't throw an overflow error (panic or
	// something else). Beside that occur an integer overflow (for that is
	// integer overflow see page 134).
	counter++
	fmt.Printf("counter = %d\n", counter)
}
