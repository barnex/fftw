package fftw


// R2CSize returns the size of the complex output array obtained by real-to-complex transform on a real input array with given size.
// A real-to-complex transform turns n0 x n1 x ... x nd reals into n0 x n1 x ... x (nd/2 + 1) complex numbers.
func R2CSize(inputSize []int)[]int{
	outputSize := make([]int, len(inputSize))
	copy(outputSize, inputSize)
	outputSize[len(outputSize)-1] =  inputSize[len(inputSize)-1]/2+1
	return outputSize
}

func checkC2CSize(n []int, howmany, in int, inembed []int, istride, idist, out int, onembed []int, ostride, odist int){
	// TODO
}

