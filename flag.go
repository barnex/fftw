package fftw

// fftw planning flags

//#include "internal/double/fftw3.h"
import "C"

import "fmt"

// FFTW planner flag:
// 	http://www.fftw.org/doc/Planner-Flags.html
type Flag uint

const (
	CONSERVE_MEMORY Flag = C.FFTW_CONSERVE_MEMORY
	DESTROY_INPUT   Flag = C.FFTW_DESTROY_INPUT  // An out-of-place transform is allowed to overwrite its input array with arbitrary data.
	ESTIMATE        Flag = C.FFTW_ESTIMATE       // Instead of time measurements, a simple heuristic is used to pick a plan quickly. The input/output arrays are not overwritten during planning.
	EXHAUSTIVE      Flag = C.FFTW_EXHAUSTIVE     // Like FFTW_PATIENT, but considers an even wider range of algorithms.
	MEASURE         Flag = C.FFTW_MEASURE        // Find an optimized plan by computing several FFTs and measuring their execution time.
	PATIENT         Flag = C.FFTW_PATIENT        // Like FFTW_MEASURE, but considers a wider range of algorithms.
	PRESERVE_INPUT  Flag = C.FFTW_PRESERVE_INPUT // An out-of-place transform must not change its input array.
	UNALIGNED       Flag = C.FFTW_UNALIGNED      // The algorithm may not impose any unusual alignment requirements on the input/output arrays.
	WISDOM_ONLY     Flag = C.FFTW_WISDOM_ONLY    // The plan is only created if wisdom is available for the given problem.
)

func (f Flag) String() string {
	if s, ok := flagName[f]; ok {
		return s
	} else {
		return fmt.Sprintf("Flag(%v)", uint(f)) // unknown flag
	}
}

// FFTW transform direction
const (
	FORWARD  = 1
	BACKWARD = -1
)

var flagName = map[Flag]string{
	CONSERVE_MEMORY: "CONSERVE_MEMORY",
	DESTROY_INPUT:   "DESTROY_INPUT",
	ESTIMATE:        "ESTIMATE",
	EXHAUSTIVE:      "EXHAUSTIVE",
	MEASURE:         "MEASURE",
	PATIENT:         "PATIENT",
	PRESERVE_INPUT:  "PRESERVE_INPUT",
	UNALIGNED:       "UNALIGNED",
	WISDOM_ONLY:     "WISDOM_ONLY",
}
