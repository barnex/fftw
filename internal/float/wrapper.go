// Internally used package float provides low-level wrappers for libfftw3f
package float

//#cgo CFLAGS: -std=gnu99 -DHAVE_CONFIG_H -fomit-frame-pointer -malign-double -fstrict-aliasing -fno-schedule-insns -ffast-math
//#cgo LDFLAGS: -lm
//#include "fftw3.h"
import "C"
import "unsafe"


// Wrapper for fftwf_plan_many_dft.
// Internal but exported for use by package fftw.
func PlanManyDft(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, sign int, flags uint) unsafe.Pointer {

	cN := to32bit(n)
	cInembed := to32bit(inembed)
	cOnembed := to32bit(onembed)

	p := C.fftwf_plan_many_dft(
		(C.int)(len(n)),
		(*C.int)(unsafe.Pointer(&cN[0])),
		(C.int)(howmany),
		(*C.fftwf_complex)(unsafe.Pointer(&in[0])),
		(*C.int)(unsafe.Pointer(&cInembed[0])),
		(C.int)(istride),
		(C.int)(idist),
		(*C.fftwf_complex)(unsafe.Pointer(&out[0])),
		(*C.int)(unsafe.Pointer(&cOnembed[0])),
		(C.int)(ostride),
		(C.int)(odist),
		(C.int)(sign),
		(C.uint)(flags))

	return unsafe.Pointer(p)
}

// Wrapper for fftwf_plan_many_dft_r2c.
// Internal but exported for use by package fftw.
func PlanManyDftR2C(n []int, howmany int, in []float32, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, flags uint) unsafe.Pointer {

	cN := to32bit(n)
	cInembed := to32bit(inembed)
	cOnembed := to32bit(onembed)

	p := C.fftwf_plan_many_dft_r2c(
		(C.int)(len(n)),
		(*C.int)(unsafe.Pointer(&cN[0])),
		(C.int)(howmany),
		(*C.float)(unsafe.Pointer(&in[0])),
		(*C.int)(unsafe.Pointer(&cInembed[0])),
		(C.int)(istride),
		(C.int)(idist),
		(*C.fftwf_complex)(unsafe.Pointer(&out[0])),
		(*C.int)(unsafe.Pointer(&cOnembed[0])),
		(C.int)(ostride),
		(C.int)(odist),
		(C.uint)(flags))

	return unsafe.Pointer(p)
}

// Wrapper for fftwf_plan_many_dft_c2r.
// Internal but exported for use by package fftw.
func PlanManyDftC2R(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []float32, onembed []int, ostride, odist int, flags uint) unsafe.Pointer {

	cN := to32bit(n)
	cInembed := to32bit(inembed)
	cOnembed := to32bit(onembed)

	p := C.fftwf_plan_many_dft_c2r(
		(C.int)(len(n)),
		(*C.int)(unsafe.Pointer(&cN[0])),
		(C.int)(howmany),
		(*C.fftwf_complex)(unsafe.Pointer(&in[0])),
		(*C.int)(unsafe.Pointer(&cInembed[0])),
		(C.int)(istride),
		(C.int)(idist),
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.int)(unsafe.Pointer(&cOnembed[0])),
		(C.int)(ostride),
		(C.int)(odist),
		(C.uint)(flags))

	return unsafe.Pointer(p)
}

// Wrapper for fftwf_execute.
// Internal but exported for use by package fftw
func Execute(plan unsafe.Pointer) {
	C.fftwf_execute(C.fftwf_plan(plan))
}

// Wrapper for fftwf_destroy_plan.
// Internal but exported for use by package fftw
func DestroyPlan(plan unsafe.Pointer) {
	C.fftwf_destroy_plan(C.fftwf_plan(plan))
}

// Wrapper for fftwf_init_threads.
// Internal but exported for use by package fftw
func InitThreads() {
	ret := C.fftwf_init_threads()
	if ret == 0 {
		panic("fftwf_init_threads failed")
	}
}

// Wrapper for fftwf_init_with_nthreads.
// Internal but exported for use by package fftw
func PlanWithNThreads(nthreads int) {
	C.fftwf_plan_with_nthreads(C.int(nthreads))
}

// C.int is 32 bit even on a 64bit system, but Go int is 32 or 64 bit.
// So we need to convert in order to pass C int arrays.
func to32bit(a []int) []int32 {
	b := make([]int32, len(a))
	for i := range b {
		b[i] = int32(a[i])
	}
	return b
}
