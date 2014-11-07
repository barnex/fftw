package fftw

// fftw planning flags

//#include "double/fftw3.h"
import "C"

import "fmt"

// FFTW planner flag:
// 	http://www.fftw.org/doc/Planner-Flags.html
type Flag uint

const (
	MEASURE         Flag = C.FFTW_MEASURE
	DESTROY_INPUT   Flag = C.FFTW_DESTROY_INPUT
	UNALIGNED       Flag = C.FFTW_UNALIGNED
	CONSERVE_MEMORY Flag = C.FFTW_CONSERVE_MEMORY
	EXHAUSTIVE      Flag = C.FFTW_EXHAUSTIVE
	PRESERVE_INPUT  Flag = C.FFTW_PRESERVE_INPUT
	PATIENT         Flag = C.FFTW_PATIENT
	ESTIMATE        Flag = C.FFTW_ESTIMATE
	WISDOM_ONLY     Flag = C.FFTW_WISDOM_ONLY
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
	MEASURE:         "MEASURE",
	DESTROY_INPUT:   "DESTROY_INPUT",
	UNALIGNED:       "UNALIGNED",
	CONSERVE_MEMORY: "CONSERVE_MEMORY",
	EXHAUSTIVE:      "EXHAUSTIVE",
	PRESERVE_INPUT:  "PRESERVE_INPUT",
	PATIENT:         "PATIENT",
	ESTIMATE:        "ESTIMATE",
	WISDOM_ONLY:     "WISDOM_ONLY"}
