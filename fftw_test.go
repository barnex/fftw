package fftw

import (
	//"testing"
	"fmt"
)

func ExamplePlanManyC2C() {
	N := 8
	data := make([]complex64, N)

	n := []int{N}
	howmany := 1
	idist := 0   // unused because howmany = 1
	odist := 0   // unused because howmany = 1
	istride := 1 // array is contiguous in memory
	ostride := 1 // array is contiguous in memory
	inembed := n
	onembed := n

	plan := PlanManyC2C(n, howmany, data, inembed, istride, idist, data, onembed, ostride, odist, FORWARD, ESTIMATE)
	defer plan.Destroy()

	data[0] = 1
	fmt.Println(data)
	plan.Execute()
	fmt.Println(data)
	plan.Execute()
	fmt.Println(data)

	// Output:
	// [(1+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i)]
	// [(1+0i) (1+0i) (1+0i) (1+0i) (1+0i) (1+0i) (1+0i) (1+0i)]
	// [(8+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i)]
}

func ExamplePlanManyZ2Z() {
	n := [2]int{2, 3}
	data := make([]complex128, n[0]*n[1])
	matrix := ReshapeZ2(data, n)

	howmany := 1
	idist := 0   // unused because howmany = 1
	odist := 0   // unused because howmany = 1
	istride := 1 // array is contiguous in memory
	ostride := 1 // array is contiguous in memory
	inembed := n[:]
	onembed := n[:]

	plan := PlanManyZ2Z(n[:], howmany, data, inembed, istride, idist, data, onembed, ostride, odist, FORWARD, ESTIMATE)
	defer plan.Destroy()

	data[0] = 1
	fmt.Println(matrix)
	plan.Execute()
	fmt.Println(matrix)
	plan.Execute()
	fmt.Println(matrix)

	// Output:
	// [[(1+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
	// [[(1+0i) (1+0i) (1+0i)] [(1+0i) (1+0i) (1+0i)]]
	// [[(6+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}
