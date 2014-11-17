package fftw

import (
	"fmt"
	"github.com/barnex/fftw/internal/double"
	"github.com/barnex/fftw/internal/float"
	"sync"
	"unsafe"
)

// Protects planners from concurrent modification.
// 	http://www.fftw.org/doc/Thread-safety.html
var lock sync.Mutex

// Plan for complex64 to complex64 FFT
type C2CPlan struct {
	handle  unsafe.Pointer // holds the C.fftwf_plan
	in, out []complex64    // pointers to data to avoid GC
}

// Wrapper for fftwf_plan_many_dft:
// 	http://www.fftw.org/doc/Advanced-Complex-DFTs.html
// Panics if the plan cannot be created. 
func PlanManyC2C(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, sign int, flags Flag) *C2CPlan {
	lock.Lock()
	defer lock.Unlock()

	p := float.PlanManyDft(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, sign:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, sign, flags))
	}
	return &C2CPlan{unsafe.Pointer(p), in, out}
}

// PlanC2C creates a complex-to-complex FFT plan of arbitrary rank. It panics when the plan can not be created. 
//
// n holds the size of the transform dimensions, len(n) is the transform's rank. 
// 
// The in, out arrays are overwritten during planning (unless FFTW_ESTIMATE is used in the flags). If in == out, the transform is in-place and the input array is overwritten. If in != out, the two arrays must not overlap (but FFTW does not check for this condition).
//
// sign can be -1 (= FFTW_FORWARD) or +1 (= FFTW_BACKWARD).
//
// flags is a bitwise OR (‘|’) of zero or more Flags.
//
// See the fftwf_plan_dft documentation:
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

// Executes the plan on the input/output arrays passed when creating the plan.
func (p *C2CPlan) Execute() {
	float.Execute(p.handle)
}

// Destroy frees the internal resources associated with this plan.
func (p *C2CPlan) Destroy() {
	lock.Lock()
	defer lock.Unlock()
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
	lock.Lock()
	defer lock.Unlock()

	p := float.PlanManyDftR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &R2CPlan{unsafe.Pointer(p), in, out}
}

// Executes the plan on the input/output arrays passed when creating the plan.
func (p *R2CPlan) Execute() {
	float.Execute(p.handle)
}

// Destroy frees the internal resources associated with this plan.
func (p *R2CPlan) Destroy() {
	lock.Lock()
	defer lock.Unlock()
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
	lock.Lock()
	defer lock.Unlock()

	p := float.PlanManyDftC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &C2RPlan{unsafe.Pointer(p), in, out}
}

// Executes the plan on the input/output arrays passed when creating the plan.
func (p *C2RPlan) Execute() {
	float.Execute(p.handle)
}

// Destroy frees the internal resources associated with this plan.
func (p *C2RPlan) Destroy() {
	lock.Lock()
	defer lock.Unlock()
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
	lock.Lock()
	defer lock.Unlock()

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

// Executes the plan on the input/output arrays passed when creating the plan.
func (p *Z2ZPlan) Execute() {
	double.Execute(p.handle)
}

// Destroy frees the internal resources associated with this plan.
func (p *Z2ZPlan) Destroy() {
	lock.Lock()
	defer lock.Unlock()
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
	lock.Lock()
	defer lock.Unlock()

	p := double.PlanManyDftR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &D2ZPlan{unsafe.Pointer(p), in, out}
}

// Executes the plan on the input/output arrays passed when creating the plan.
func (p *D2ZPlan) Execute() {
	double.Execute(p.handle)
}

// Destroy frees the internal resources associated with this plan.
func (p *D2ZPlan) Destroy() {
	lock.Lock()
	defer lock.Unlock()
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
	lock.Lock()
	defer lock.Unlock()

	p := double.PlanManyDftC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, flags))
	}
	return &Z2DPlan{unsafe.Pointer(p), in, out}
}

// Executes the plan on the input/output arrays passed when creating the plan.
func (p *Z2DPlan) Execute() {
	double.Execute(p.handle)
}

// Destroy frees the internal resources associated with this plan.
func (p *Z2DPlan) Destroy() {
	lock.Lock()
	defer lock.Unlock()
	double.DestroyPlan(p.handle)
	p.handle = nil
	p.in = nil
	p.out = nil
}
