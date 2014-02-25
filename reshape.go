package fftw

// utilities to construct multi-dimensional arrays backed by contiguous storage

import "fmt"

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeR2(array []float32, N [2]int) [][]float32 {
	return ReshapeR3(array, [3]int{1, N[0], N[1]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeR3(array []float32, N [3]int) [][][]float32 {
	return ReshapeR4(array, [4]int{1, N[0], N[1], N[2]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeR4(array []float32, N [4]int) [][][][]float32 {
	if prod(N[:]) != len(array) {
		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
	}
	sliced := make([][][][]float32, N[0])
	for i := range sliced {
		sliced[i] = make([][][]float32, N[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]float32, N[2])
		}
	}
	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
			}
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeC2(array []complex64, N [2]int) [][]complex64 {
	return ReshapeC3(array, [3]int{1, N[0], N[1]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeC3(array []complex64, N [3]int) [][][]complex64 {
	return ReshapeC4(array, [4]int{1, N[0], N[1], N[2]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeC4(array []complex64, N [4]int) [][][][]complex64 {
	if prod(N[:]) != len(array) {
		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
	}
	sliced := make([][][][]complex64, N[0])
	for i := range sliced {
		sliced[i] = make([][][]complex64, N[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]complex64, N[2])
		}
	}
	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
			}
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeD2(array []float64, N [2]int) [][]float64 {
	return ReshapeD3(array, [3]int{1, N[0], N[1]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeD3(array []float64, N [3]int) [][][]float64 {
	return ReshapeD4(array, [4]int{1, N[0], N[1], N[2]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeD4(array []float64, N [4]int) [][][][]float64 {
	if prod(N[:]) != len(array) {
		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
	}
	sliced := make([][][][]float64, N[0])
	for i := range sliced {
		sliced[i] = make([][][]float64, N[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]float64, N[2])
		}
	}
	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
			}
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeZ2(array []complex128, N [2]int) [][]complex128 {
	return ReshapeZ3(array, [3]int{1, N[0], N[1]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeZ3(array []complex128, N [3]int) [][][]complex128 {
	return ReshapeZ4(array, [4]int{1, N[0], N[1], N[2]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func ReshapeZ4(array []complex128, N [4]int) [][][][]complex128 {
	if prod(N[:]) != len(array) {
		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
	}
	sliced := make([][][][]complex128, N[0])
	for i := range sliced {
		sliced[i] = make([][][]complex128, N[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]complex128, N[2])
		}
	}
	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
			}
		}
	}
	return sliced
}

func prod(s []int) int {
	p := 1
	for _, s := range s {
		p *= s
	}
	return p
}
