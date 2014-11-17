#FFTW3 Go interface [![GoDoc](https://godoc.org/github.com/barnex/fftw?status.svg)](https://godoc.org/github.com/barnex/fftw)   [![Build Status](https://travis-ci.org/barnex/fftw.svg)](https://travis-ci.org/barnex/fftw)

Go wrapper for FFTW3.3.4, without dependencies. FFTW's C code is embedded in this package, resulting binaries are statically linked.
You can just `go get -x github.com/barnex/fftw` (-x to show what it's doing, compilation takes long due to C files).


## Copyright

This package inherits the GNU General Public License from FFTW, which carries this notice:

FFTW is Copyright © 2003, 2007-11 Matteo Frigo, Copyright © 2003, 2007-11 Massachusetts Institute of Technology.

FFTW is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later version.


