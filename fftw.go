package fftw

import (
	"fmt"
	"github.com/barnex/fftw/float"
	"github.com/barnex/fftw/double"
	"unsafe"
)

// Plan for complex64 to complex64 FFT
type PlanC2C struct {
	handle  unsafe.Pointer // holds the C.fftwf_plan
	in, out []complex64 // pointers to data to avoid GC
}

// Wrapper for fftwf_plan_many_dft:
// 	http://www.fftw.org/doc/Advanced-Complex-DFTs.html#Advanced-Complex-DFTs
func PlanManyC2C(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, sign int, flags Flag) *PlanC2C {

	p := float.PlanManyDft(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, sign:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, sign, flags))
	}
	return &PlanC2C{unsafe.Pointer(p), in, out}
}

func(p*PlanC2C)Execute(){
	float.Execute(p.handle)
}

func(p*PlanC2C)Destroy(){
	float.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

// Plan for float32 to complex64 FFT
type PlanR2C struct {
	handle unsafe.Pointer
	in     []float32
	out    []complex64
}

// Wrapper for fftwf_plan_many_dft_r2c:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html#Advanced-Real_002ddata-DFTs
func PlanManyR2C(n []int, howmany int, in []float32, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, flags Flag) *PlanR2C {

	p := float.PlanManyDftR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &PlanR2C{unsafe.Pointer(p), in, out}
}


func(p*PlanR2C)Execute(){
	float.Execute(p.handle)
}

func(p*PlanR2C)Destroy(){
	float.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

type PlanC2R struct {
	handle unsafe.Pointer
	in     []complex64
	out    []float32
}

// Wrapper for fftwf_plan_many_dft_c2r:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html#Advanced-Real_002ddata-DFTs
func PlanManyC2R(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []float32, onembed []int, ostride, odist int, flags Flag) *PlanC2R {

	p := float.PlanManyDftC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &PlanC2R{unsafe.Pointer(p), in, out}
}


func(p*PlanC2R)Execute(){
	float.Execute(p.handle)
}

func(p*PlanC2R)Destroy(){
	float.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}














// Plan for complex128 to complex128 FFT
type PlanZ2Z struct {
	handle  unsafe.Pointer // holds the C.fftwf_plan
	in, out []complex128 // pointers to data to avoid GC
}

// Wrapper for fftw_plan_many_dft:
// 	http://www.fftw.org/doc/Advanced-Complex-DFTs.html#Advanced-Complex-DFTs
func PlanManyZ2Z(n []int, howmany int, in []complex128, inembed []int, istride, idist int,
	out []complex128, onembed []int, ostride, odist int, sign int, flags Flag) *PlanZ2Z {

	p := double.PlanManyDft(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, sign:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, sign, flags))
	}
	return &PlanZ2Z{unsafe.Pointer(p), in, out}
}

func(p*PlanZ2Z)Execute(){
	double.Execute(p.handle)
}

func(p*PlanZ2Z)Destroy(){
	double.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

// Plan for float64 to complex128 FFT
type PlanD2Z struct {
	handle unsafe.Pointer
	in     []float64
	out    []complex128
}

// Wrapper for fftw_plan_many_dft_r2c:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html#Advanced-Real_002ddata-DFTs
func PlanManyD2Z(n []int, howmany int, in []float64, inembed []int, istride, idist int,
	out []complex128, onembed []int, ostride, odist int, flags Flag) *PlanD2Z {

	p := double.PlanManyDftR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &PlanD2Z{unsafe.Pointer(p), in, out}
}


func(p*PlanD2Z)Execute(){
	double.Execute(p.handle)
}

func(p*PlanD2Z)Destroy(){
	double.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

type PlanZ2D struct {
	handle unsafe.Pointer
	in     []complex128
	out    []float64
}

// Wrapper for fftw_plan_many_dft_c2r:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html#Advanced-Real_002ddata-DFTs
func PlanManyZ2D(n []int, howmany int, in []complex128, inembed []int, istride, idist int,
	out []float64, onembed []int, ostride, odist int, flags Flag) *PlanZ2D {

	p := double.PlanManyDftC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &PlanZ2D{unsafe.Pointer(p), in, out}
}


func(p*PlanZ2D)Execute(){
	double.Execute(p.handle)
}

func(p*PlanZ2D)Destroy(){
	double.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

