package fftw

import (
	"unsafe"
	"github.com/barnex/fftw/internal/float"
	"github.com/barnex/fftw/internal/double"
)

// A Plan contains all information necessary to compute the transform, 
// including the input and output arrays.
type Plan interface{
 	Execute() // Executes the plan on the input/output arrays passed when creating the plan.
	Destroy() // Frees the internal resources associated with this plan (not the input/output arrays).
}

type floatHandle struct{
	handle unsafe.Pointer // holds the C.fftwf_plan
}

func(h floatHandle)Execute(){
	float.Execute(h.handle)
}


func(h *floatHandle)Destroy(){
	lock.Lock()
	defer lock.Unlock()
	float.DestroyPlan(h.handle)
	h.handle = nil
}

type doubleHandle struct{
	handle unsafe.Pointer
}


func(h doubleHandle)Execute(){
	double.Execute(h.handle)
}


func(h *doubleHandle)Destroy(){
	lock.Lock()
	defer lock.Unlock()
	double.DestroyPlan(h.handle)
	h.handle = nil
}
