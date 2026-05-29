// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

func checkgoarm() {
	// SotaConnect legacy iOS targets GOARM=7 armv7 devices. Modern Go no
	// longer provides getncpu for this removed port, so the old multi-CPU
	// GOARM guard is intentionally disabled in this local backport.
}

//go:nosplit
func cputicks() int64 {
	// Currently cputicks() is used in blocking profiler and to seed runtime·fastrand().
	// runtime·nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	return nanotime()
}
