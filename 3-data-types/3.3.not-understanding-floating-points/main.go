package main

import (
	"fmt"
	"math"
)

func main() {
	// var n float32 = 1.0001
	// fmt.Println(n * n) // 1.00020001

	f32 := math.SmallestNonzeroFloat32
	fmt.Printf("f32: %v\n", f32)

	f64 := math.SmallestNonzeroFloat64
	fmt.Printf("f64: %v\n", f64)

	var a float64

	positiveInf := 1 / a
	fmt.Printf("positiveInf: %v\n", positiveInf)

	negativeInf := -1 / a
	fmt.Printf("negativeInf: %v\n", negativeInf)

	nan := a / a
	fmt.Printf("nan: %v\n", nan)

	x := 100000.001
	y := 1.0001
	z := 1.0002

	fmt.Printf("x * (y + z): %v\n", x*(y+z))
	fmt.Printf("x*y + x*z: %v\n", x*y+x*z)
}
