package fftw

import (
	"github.com/barnex/fftw/internal/double"
	"github.com/barnex/fftw/internal/float"
)

// InitThreads should be called once to initialize multi-threaded FFTW. See:
// 	http://www.fftw.org/doc/Usage-of-Multi_002dthreaded-FFTW.html
func InitThreads() {
	float.InitThreads()
	double.InitThreads()
}

// PlanWithNThreads(n) causes all subsequent plans to use at most n threads. See:
// 	http://www.fftw.org/doc/Usage-of-Multi_002dthreaded-FFTW.html
func PlanWithNThreads(nthreads int) {
	float.PlanWithNThreads(nthreads)
	double.PlanWithNThreads(nthreads)
}
