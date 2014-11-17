#FFTW3 Go interface [![GoDoc](https://godoc.org/github.com/barnex/fftw?status.svg)](https://godoc.org/github.com/barnex/fftw)   [![Build Status](https://travis-ci.org/barnex/fftw.svg)](https://travis-ci.org/barnex/fftw)

Go wrapper for FFTW3.3.4, without dependencies. FFTW's C code is embedded in this package, resulting binaries are statically linked.
You can just `go get -x github.com/barnex/fftw` (-x to show what it's doing, compilation takes long due to C files).

Single and doulbe precission real-to-complex, complex-to-real and complex-to-complex transforms are provided for arbitrary rank and dimensions.


## Using Plans

To perform an FFT, one first creates a Plan for the given dimensions and input/output arrays. Once a plan is created, it can be executed rapidly and as many times as desired. E.g.:

```go
	N := 8
	data := make([]complex64, N)
	plan := PlanC2C([]int{N}, data, data, FORWARD, ESTIMATE)

	// data = ...
	plan.Execute()
	// use data
```


## Data storage

This package acts on contiguous arrays, even for multi-dimensional input. The package
 	`github.com/barnex/matrix`
may be used to easily construct and access multi-dimensional arrays with contiguous underlying storage.



## Copyright

This package inherits the GNU General Public License from FFTW, which carries this notice:

FFTW is Copyright © 2003, 2007-11 Matteo Frigo, Copyright © 2003, 2007-11 Massachusetts Institute of Technology.

FFTW is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later version.


