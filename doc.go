/*

Go wrapper for FFTW3.3.4, without dependencies. FFTW's C code is embedded in this package, resulting binaries are statically linked.
 
Single and doulbe precission real-to-complex, complex-to-real and complex-to-complex transforms are provided for arbitrary rank and dimensions.


Example

Example of a 1-dimensional complex-to-complex transform on 8 elements, in-place:

	N := 8
	data := make([]complex64, N)
	plan := PlanC2C([]int{N}, data, data, FORWARD, ESTIMATE)

	data[0] = 1
	plan.Execute()
	fmt.Println(data)


License


FFTW is Copyright © 2003, 2007-11 Matteo Frigo, Copyright © 2003, 2007-11 Massachusetts Institute of Technology.

FFTW is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA You can also find the GPL on the GNU web site.

In addition, we kindly ask you to acknowledge FFTW and its authors in any program or publication in which you use FFTW. (You are not required to do so; it is up to your common sense to decide whether you want to comply with this request or not.) For general publications, we suggest referencing: Matteo Frigo and Steven G. Johnson, “The design and implementation of FFTW3,” Proc. IEEE 93 (2), 216–231 (2005).

Non-free versions of FFTW are available under terms different from those of the General Public License. (e.g. they do not require you to accompany any object code using FFTW with the corresponding source code.) For these alternative terms you must purchase a license from MIT's Technology Licensing Office. Users interested in such a license should contact us (fftw@fftw.org) for more information.

	http://www.fftw.org/doc/

Go wrappers and utility functions (Cast, Reshape) written by Arne Vansteenkiste <arne@vansteenkiste.io>, 2014.
*/
package fftw
