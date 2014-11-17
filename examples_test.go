package fftw

import (
	"fmt"
	"github.com/barnex/matrix"
)

func ExamplePlanC2C() {
	N := 8
	data := make([]complex64, N)

	n := []int{N}
	plan := PlanC2C(n, data, data, FORWARD, ESTIMATE)
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


func ExamplePlanR2C() {
	N := 5
	in := make([]float32, N)
	out := make([]complex64, N/2+1)

	n := []int{N}
	plan := PlanR2C(n, in, out, ESTIMATE)
	defer plan.Destroy()

	in[0] = 1
	fmt.Println(in)
	plan.Execute()
	fmt.Println(out)

	// Output:
	// [1 0 0 0 0]
	// [(1+0i) (1+0i) (1+0i)]
}


func ExamplePlanC2R() {
	N := 5
	in := make([]complex64, N/2+1)
	out := make([]float32, N)

	n := []int{N}
	plan := PlanC2R(n, in, out, ESTIMATE)
	defer plan.Destroy()

	in[0] = 1
	fmt.Println(in)
	plan.Execute()
	fmt.Println(out)

	// Output:
	// [(1+0i) (0+0i) (0+0i)]
	// [1 1 1 1 1]
}



func ExamplePlanD2Z() {
	N := 4
	in := make([]float64, N)
	out := make([]complex128, N/2+1)

	n := []int{N}
	plan := PlanD2Z(n, in, out, ESTIMATE)
	defer plan.Destroy()

	in[0] = 1
	fmt.Println(in)
	plan.Execute()
	fmt.Println(out)

	// Output:
	// [1 0 0 0]
	// [(1+0i) (1+0i) (1+0i)]
}


func ExamplePlanZ2D() {
	N := 4
	in := make([]complex128, N/2+1)
	out := make([]float64, N)

	n := []int{N}
	plan := PlanZ2D(n, in, out, ESTIMATE)
	defer plan.Destroy()

	in[0] = 1
	fmt.Println(in)
	plan.Execute()
	fmt.Println(out)

	// Output:
	// [(1+0i) (0+0i) (0+0i)]
	// [1 1 1 1]
}



func ExamplePlanZ2Z() {
	N := 8
	data := make([]complex128, N)

	n := []int{N}
	plan := PlanZ2Z(n, data, data, FORWARD, ESTIMATE)
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
	matrix := matrix.ReshapeZ2(data, n)

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
