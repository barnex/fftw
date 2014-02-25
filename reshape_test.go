package fftw

import "fmt"

func ExampleReshapeR2() {
	array := []float32{0, 1, 2, 3, 4, 5}
	matrix := ReshapeR2(array, [2]int{2, 3})
	fmt.Println(matrix)

	// Output:
	// [[0 1 2] [3 4 5]]
}
