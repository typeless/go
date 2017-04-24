package hal

import "unsafe"

// MemclrNoHeapPointers clears n bytes starting at ptr.
//
// Usually you should use typedmemclr. memclrNoHeapPointers should be
// used only when the caller knows that *ptr contains no heap pointers
// because either:
//
// 1. *ptr is initialized memory and its type is pointer-free.
//
// 2. *ptr is uninitialized memory (e.g., memory that's being reused
//    for a new allocation) and hence contains only "junk".
//
// in memclr_*.s
//go:noescape
func MemclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
