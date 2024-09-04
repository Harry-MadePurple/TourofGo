package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	difference := 0.0
	// part 1
	//for i := 0; i < 100; i++ {
	//	z -= (z*z - x) / (2 * z)
	//	fmt.Println("z: ", z)
	//}
	// part 2
	for {

		z, difference = z-(z*z-x)/(2*z), z

		if math.Abs(difference-z) < 0.0000000000001 {
			break
		}
		fmt.Println("z: ", z)

	}
	return z
}

func main() {
	square := 3.0
	actual := Sqrt(square)
	calculated := math.Sqrt(square)
	fmt.Println("Actual: ", actual)
	fmt.Println("Calculated: ", calculated)
	fmt.Println("Difference: ", math.Abs(actual-calculated))
}
