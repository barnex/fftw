#! /bin/bash

set -e

#./configure --enable-float --enable-threads --with-combined-threads --disable-fortran --enable-sse --enable-sse2 --enable-avx
#make -j 8

dirs='kernel simd-support dft dft/scalar dft/scalar/codelets dft/simd dft/simd/common dft/simd/sse2 rdft rdft/scalar rdft/scalar/r2cf rdft/scalar/r2cb rdft/scalar/r2r rdft/simd rdft/simd/common rdft/simd/sse2 reodft api threads'

rm -rf all
mkdir all


for d in $dirs; do
	(cp -v $d/*.h all || echo "")
	prefix=$(echo $d | sed 's-/-_-g')
	for c in $d/*.c; do
		(cp -v $c all/$prefix'_'$(basename $c) || echo "")
	done;
done;

cp config.h all;

cd all



for f in dft_simd_sse2*.c; do
	echo $f
	echo '#define SIMD_HEADER "simd-sse2.h"' > $f
	cat $(echo $f | sed 's#dft_simd_sse2#dft_simd_common#g') >> $f;
done




for f in rdft_simd_sse2*.c; do
	echo $f
	echo '#define SIMD_HEADER "simd-sse2.h"' > $f
	cat $(echo $f | sed 's#rdft_simd_sse2#rdft_simd_common#g') >> $f;
done

rm dft_simd_common_*.c
rm rdft_simd_common_*.c
rm threads_openmp.c

#gcc -std=gnu99 -DHAVE_CONFIG_H -O3 -fomit-frame-pointer -mtune=native -malign-double -fstrict-aliasing -fno-schedule-insns -ffast-math -c *.c

ar rcs libfftw3.a *.o
