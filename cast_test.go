package fftw

import "fmt"

func ExampleCastCtoR() {
	cmplx := []complex64{complex(1, 2), complex(3, 4)}
	real := CastCtoR(cmplx)
	fmt.Println("CastCtoR(", cmplx, ") =", real)
	fmt.Println("CastRtoC(", real, ") =", CastRtoC(real))

	// Output:
	// CastCtoR( [(1+2i) (3+4i)] ) = [1 2 3 4]
	// CastRtoC( [1 2 3 4] ) = [(1+2i) (3+4i)]
}
