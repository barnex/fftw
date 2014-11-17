package fftw

import (
	"fmt"
	"unsafe"

	"github.com/barnex/fftw/internal/double"
	"github.com/barnex/fftw/internal/float"
)

// An FFT Plan is created once to set up an FFT with certain size and input/output arrays.
// After plan creation, it can be executed quickly and as many times as desired.
// Creating multiple plans with the same size is cheap after the first one.
type Plan interface {
	Execute() // Executes the plan on the input/output arrays passed when creating the plan.
	Destroy() // Frees the internal resources associated with this plan (not the input/output arrays).
}

// single precission plan
type floatHandle struct {
	handle  unsafe.Pointer // holds the C.fftwf_plan
	in, out unsafe.Pointer
}

func (h *floatHandle) Execute() {
	float.Execute(h.handle)
}

func (h *floatHandle) Destroy() {
	lock.Lock()
	defer lock.Unlock()
	float.DestroyPlan(h.handle)
	h.handle = nil
	h.in = nil
	h.out = nil
}

// double precission plan
type doubleHandle struct {
	handle  unsafe.Pointer
	in, out unsafe.Pointer
}

func (h doubleHandle) Execute() {
	double.Execute(h.handle)
}

func (h *doubleHandle) Destroy() {
	lock.Lock()
	defer lock.Unlock()
	double.DestroyPlan(h.handle)
	h.handle = nil
	h.in = nil
	h.out = nil
}

// panic if p == nil
func checkPlan(p unsafe.Pointer, n []int, howmany int, inembed []int, istride, idist int, onembed []int, ostride, odist int, sign int, flags uint) {
	if p == nil {
		panic(fmt.Errorf("invalid plan: n:%v, howmany:%v, inembed:%v, istride:%v, idist:%v, onembed:%v, ostride:%v, odist:%v, sign:%v, flags:%v", n, howmany, inembed, istride, idist, onembed, ostride, odist, sign, flags))
	}
}
