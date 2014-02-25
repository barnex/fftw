package fftw

import (
	"fmt"
	"github.com/barnex/fftw/float"
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
