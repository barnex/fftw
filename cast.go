package fftw

// functions for casting between real/complex arrays

import (
	"fmt"
	"unsafe"
)

// Re-interprets the slice of complex numbers
// as a slice of real numbers, twice as long.
// Underlying storage is shared.
func CastCtoR(c []complex64) []float32 {
	return (*(*[1<<31 - 1]float32)(unsafe.Pointer(&c[0])))[:2*len(c)]
}

// Re-interprets the slice of real numbers
// as a slice of complex numbers, half as long.
// Underlying storage is shared.
func CastRtoC(r []float32) []complex64 {
	if len(r)%2 != 0 {
		panic(fmt.Errorf("input len should be even, have: %v", len(r)))
	}
	return (*(*[1<<31 - 1]complex64)(unsafe.Pointer(&r[0])))[:len(r)/2]
}

// Re-interprets the slice of complex numbers
// as a slice of real numbers, twice as long.
// Underlying storage is shared.
func CastZtoD(c []complex128) []float64 {
	return (*(*[1<<31 - 1]float64)(unsafe.Pointer(&c[0])))[:2*len(c)]
}

// Re-interprets the slice of real numbers
// as a slice of complex numbers, half as long.
// Underlying storage is shared.
func CastDtoZ(r []float64) []complex128 {
	if len(r)%2 != 0 {
		panic(fmt.Errorf("input len should be even, have: %v", len(r)))
	}
	return (*(*[1<<31 - 1]complex128)(unsafe.Pointer(&r[0])))[:len(r)/2]
}
