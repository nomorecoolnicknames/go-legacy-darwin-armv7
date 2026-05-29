// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT _rt0_arm_darwin(SB),NOSPLIT|NOFRAME,$0
	MOVW	(R13), R0	// argc
	MOVW	$4(R13), R1	// argv
	B	runtime·rt0_go(SB)

TEXT _rt0_arm_darwin_lib(SB),NOSPLIT,$0
	B	_rt0_arm_lib(SB)
