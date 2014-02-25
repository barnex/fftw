package fftw

import (
	"fmt"
	"github.com/barnex/fftw/double"
	"github.com/barnex/fftw/float"
	"unsafe"
)

// Plan for complex64 to complex64 FFT
type C2CPlan struct {
	handle  unsafe.Pointer // holds the C.fftwf_plan
	in, out []complex64    // pointers to data to avoid GC
}

// Wrapper for fftwf_plan_many_dft:
// 	http://www.fftw.org/doc/Advanced-Complex-DFTs.html
func PlanManyC2C(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, sign int, flags Flag) *C2CPlan {

	p := float.PlanManyDft(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, sign:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, sign, flags))
	}
	return &C2CPlan{unsafe.Pointer(p), in, out}
}

// Provides the functionality of fftwf_plan_dft:
// 	http://www.fftw.org/doc/Complex-DFTs.html
func PlanC2C(n []int, in []complex64, out []complex64, sign int, flags Flag) *C2CPlan {
	howmany := 1
	idist := 0
	odist := 0
	istride := 1
	ostride := 1
	inembed := n
	onembed := n
	return PlanManyC2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, flags)
}

func (p *C2CPlan) Execute() {
	float.Execute(p.handle)
}

func (p *C2CPlan) Destroy() {
	float.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

// Plan for float32 to complex64 FFT
type R2CPlan struct {
	handle unsafe.Pointer
	in     []float32
	out    []complex64
}

// Wrapper for fftwf_plan_many_dft_r2c:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html
func PlanManyR2C(n []int, howmany int, in []float32, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, flags Flag) *R2CPlan {

	p := float.PlanManyDftR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &R2CPlan{unsafe.Pointer(p), in, out}
}

func (p *R2CPlan) Execute() {
	float.Execute(p.handle)
}

func (p *R2CPlan) Destroy() {
	float.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

type C2RPlan struct {
	handle unsafe.Pointer
	in     []complex64
	out    []float32
}

// Wrapper for fftwf_plan_many_dft_c2r:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html
func PlanManyC2R(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []float32, onembed []int, ostride, odist int, flags Flag) *C2RPlan {

	p := float.PlanManyDftC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &C2RPlan{unsafe.Pointer(p), in, out}
}

func (p *C2RPlan) Execute() {
	float.Execute(p.handle)
}

func (p *C2RPlan) Destroy() {
	float.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

// Plan for complex128 to complex128 FFT
type Z2ZPlan struct {
	handle  unsafe.Pointer // holds the C.fftwf_plan
	in, out []complex128   // pointers to data to avoid GC
}

// Wrapper for fftw_plan_many_dft:
// 	http://www.fftw.org/doc/Advanced-Complex-DFTs.html
func PlanManyZ2Z(n []int, howmany int, in []complex128, inembed []int, istride, idist int,
	out []complex128, onembed []int, ostride, odist int, sign int, flags Flag) *Z2ZPlan {

	p := double.PlanManyDft(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, sign:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, sign, flags))
	}
	return &Z2ZPlan{unsafe.Pointer(p), in, out}
}

// Provides the functionality of fftw_plan_dft:
// 	http://www.fftw.org/doc/Complex-DFTs.html
func PlanZ2Z(n []int, in []complex128, out []complex128, sign int, flags Flag) *Z2ZPlan {
	howmany := 1
	idist := 0
	odist := 0
	istride := 1
	ostride := 1
	inembed := n
	onembed := n
	return PlanManyZ2Z(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, flags)
}

func (p *Z2ZPlan) Execute() {
	double.Execute(p.handle)
}

func (p *Z2ZPlan) Destroy() {
	double.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

// Plan for float64 to complex128 FFT
type D2ZPlan struct {
	handle unsafe.Pointer
	in     []float64
	out    []complex128
}

// Wrapper for fftw_plan_many_dft_r2c:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html
func PlanManyD2Z(n []int, howmany int, in []float64, inembed []int, istride, idist int,
	out []complex128, onembed []int, ostride, odist int, flags Flag) *D2ZPlan {

	p := double.PlanManyDftR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &D2ZPlan{unsafe.Pointer(p), in, out}
}

func (p *D2ZPlan) Execute() {
	double.Execute(p.handle)
}

func (p *D2ZPlan) Destroy() {
	double.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}

type Z2DPlan struct {
	handle unsafe.Pointer
	in     []complex128
	out    []float64
}

// Wrapper for fftw_plan_many_dft_c2r:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html
func PlanManyZ2D(n []int, howmany int, in []complex128, inembed []int, istride, idist int,
	out []float64, onembed []int, ostride, odist int, flags Flag) *Z2DPlan {

	p := double.PlanManyDftC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &Z2DPlan{unsafe.Pointer(p), in, out}
}

func (p *Z2DPlan) Execute() {
	double.Execute(p.handle)
}

func (p *Z2DPlan) Destroy() {
	double.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}
