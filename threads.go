package fftw

import(
	"github.com/barnex/fftw/double"
	"github.com/barnex/fftw/float"
)

func InitThreads() {
	float.InitThreads()
	double.InitThreads()
}

func PlanWithNThreads(nthreads int) {
	float.PlanWithNThreads(nthreads)
	double.PlanWithNThreads(nthreads)
}
