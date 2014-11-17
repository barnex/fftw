package fftw

import (
	"sync"
	"unsafe"

	"github.com/barnex/fftw/internal/double"
	"github.com/barnex/fftw/internal/float"
)

// Protects planners from concurrent modification.
// 	http://www.fftw.org/doc/Thread-safety.html
var lock sync.Mutex

// Wrapper for fftwf_plan_many_dft:
// 	http://www.fftw.org/doc/Advanced-Complex-DFTs.html
// Panics if the plan cannot be created.
func PlanManyC2C(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, sign int, flags Flag) Plan {
	lock.Lock()
	defer lock.Unlock()

	p := float.PlanManyDft(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, uint(flags))
	checkPlan(p, n, howmany, inembed, istride, idist, onembed, ostride, odist, sign, uint(flags))
	checkC2CSize(n, howmany, len(in), inembed, istride, idist, len(out), onembed, ostride, odist)
	return &floatHandle{p, unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0])}
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
func PlanC2C(n []int, in []complex64, out []complex64, sign int, flags Flag) Plan {
	howmany := 1
	idist := 0
	odist := 0
	istride := 1
	ostride := 1
	inembed := n
	onembed := n
	return PlanManyC2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, flags)
}

// Wrapper for fftwf_plan_many_dft_r2c:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html
func PlanManyR2C(n []int, howmany int, in []float32, inembed []int, istride, idist int,
	out []complex64, onembed []int, ostride, odist int, flags Flag) Plan {
	lock.Lock()
	defer lock.Unlock()

	p := float.PlanManyDftR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	checkPlan(p, n, howmany, inembed, istride, idist, onembed, ostride, odist, FORWARD, uint(flags))
	return &floatHandle{p, unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0])}
}


// PlanR2C creates a real-to-complex FFT plan of arbitrary rank. It panics when the plan can not be created.
// The size of the output array is roughly half the size of the input array, see R2CSize.
//
// n holds the size of the transform dimensions, len(n) is the transform's rank.
//
// The in, out arrays are overwritten during planning (unless FFTW_ESTIMATE is used in the flags). If in == out, the transform is in-place and the input array is overwritten. If in != out, the two arrays must not overlap (but FFTW does not check for this condition).
//
// flags is a bitwise OR (‘|’) of zero or more Flags.
//
// See the fftw_plan_dft_r2c documentation:
// 	http://www.fftw.org/doc/Real_002ddata-DFTs.html
func PlanR2C(n []int, in []float32, out []complex64, flags Flag) Plan {
	howmany := 1
	idist := 0
	odist := 0
	istride := 1
	ostride := 1
	inembed := n
	onembed := n
	return PlanManyR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, flags)
}


// PlanC2R creates a complex-to-real FFT plan of arbitrary rank. It panics when the plan can not be created.
//
// n holds the size of the transform dimensions, len(n) is the transform's rank.
//
// The in, out arrays are overwritten during planning (unless FFTW_ESTIMATE is used in the flags). If in == out, the transform is in-place and the input array is overwritten. If in != out, the two arrays must not overlap (but FFTW does not check for this condition).
//
// flags is a bitwise OR (‘|’) of zero or more Flags.
//
// See the fftw_plan_dft_r2c documentation:
// 	http://www.fftw.org/doc/Real_002ddata-DFTs.html
func PlanC2R(n []int, in []complex64, out []float32, flags Flag) Plan {
	howmany := 1
	idist := 0
	odist := 0
	istride := 1
	ostride := 1
	inembed := n
	onembed := n
	return PlanManyC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, flags)
}


// Wrapper for fftwf_plan_many_dft_c2r:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html
func PlanManyC2R(n []int, howmany int, in []complex64, inembed []int, istride, idist int,
	out []float32, onembed []int, ostride, odist int, flags Flag) Plan {
	lock.Lock()
	defer lock.Unlock()

	p := float.PlanManyDftC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	checkPlan(p, n, howmany, inembed, istride, idist, onembed, ostride, odist, BACKWARD, uint(flags))
	return &floatHandle{p, unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0])}
}

// PlanZ2Z creates a complex-to-complex FFT plan of arbitrary rank. It panics when the plan can not be created.
//
// n holds the size of the transform dimensions, len(n) is the transform's rank.
//
// The in, out arrays are overwritten during planning (unless FFTW_ESTIMATE is used in the flags). If in == out, the transform is in-place and the input array is overwritten. If in != out, the two arrays must not overlap (but FFTW does not check for this condition).
//
// sign can be -1 (= FFTW_FORWARD) or +1 (= FFTW_BACKWARD).
//
// flags is a bitwise OR (‘|’) of zero or more Flags.
//
// See the fftw_plan_dft documentation:
// 	http://www.fftw.org/doc/Complex-DFTs.html
func PlanManyZ2Z(n []int, howmany int, in []complex128, inembed []int, istride, idist int,
	out []complex128, onembed []int, ostride, odist int, sign int, flags Flag) Plan {
	lock.Lock()
	defer lock.Unlock()

	p := double.PlanManyDft(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, uint(flags))
	checkPlan(p, n, howmany, inembed, istride, idist, onembed, ostride, odist, sign, uint(flags))
	return &doubleHandle{p, unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0])}
}

// Provides the functionality of fftw_plan_dft:
// 	http://www.fftw.org/doc/Complex-DFTs.html
func PlanZ2Z(n []int, in []complex128, out []complex128, sign int, flags Flag) Plan {
	howmany := 1
	idist := 0
	odist := 0
	istride := 1
	ostride := 1
	inembed := n
	onembed := n
	return PlanManyZ2Z(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, sign, flags)
}

// Wrapper for fftw_plan_many_dft_r2c:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html
func PlanManyD2Z(n []int, howmany int, in []float64, inembed []int, istride, idist int,
	out []complex128, onembed []int, ostride, odist int, flags Flag) Plan {
	lock.Lock()
	defer lock.Unlock()

	p := double.PlanManyDftR2C(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	checkPlan(p, n, howmany, inembed, istride, idist, onembed, ostride, odist, FORWARD, uint(flags))
	return &doubleHandle{p, unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0])}
}

// Wrapper for fftw_plan_many_dft_c2r:
// 	http://www.fftw.org/doc/Advanced-Real_002ddata-DFTs.html
func PlanManyZ2D(n []int, howmany int, in []complex128, inembed []int, istride, idist int,
	out []float64, onembed []int, ostride, odist int, flags Flag) Plan {
	lock.Lock()
	defer lock.Unlock()

	p := double.PlanManyDftC2R(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, uint(flags))
	checkPlan(p, n, howmany, inembed, istride, idist, onembed, ostride, odist, BACKWARD, uint(flags))
	return &doubleHandle{p, unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0])}
}


// PlanD2Z creates a real-to-complex FFT plan of arbitrary rank. It panics when the plan can not be created.
// The size of the output array is roughly half the size of the input array, see R2CSize.
//
// n holds the size of the transform dimensions, len(n) is the transform's rank.
//
// The in, out arrays are overwritten during planning (unless FFTW_ESTIMATE is used in the flags). If in == out, the transform is in-place and the input array is overwritten. If in != out, the two arrays must not overlap (but FFTW does not check for this condition).
//
// flags is a bitwise OR (‘|’) of zero or more Flags.
//
// See the fftw_plan_dft_r2c documentation:
// 	http://www.fftw.org/doc/Real_002ddata-DFTs.html
func PlanD2Z(n []int, in []float64, out []complex128, flags Flag) Plan {
	howmany := 1
	idist := 0
	odist := 0
	istride := 1
	ostride := 1
	inembed := n
	onembed := n
	return PlanManyD2Z(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, flags)
}


// PlanZ2D creates a complex-to-real FFT plan of arbitrary rank. It panics when the plan can not be created.
//
// n holds the size of the transform dimensions, len(n) is the transform's rank.
//
// The in, out arrays are overwritten during planning (unless FFTW_ESTIMATE is used in the flags). If in == out, the transform is in-place and the input array is overwritten. If in != out, the two arrays must not overlap (but FFTW does not check for this condition).
//
// flags is a bitwise OR (‘|’) of zero or more Flags.
//
// See the fftw_plan_dft_r2c documentation:
// 	http://www.fftw.org/doc/Real_002ddata-DFTs.html
func PlanZ2D(n []int, in []complex128, out []float64, flags Flag) Plan {
	howmany := 1
	idist := 0
	odist := 0
	istride := 1
	ostride := 1
	inembed := n
	onembed := n
	return PlanManyZ2D(n, howmany, in, inembed, istride, idist, out, onembed, ostride, odist, flags)
}

