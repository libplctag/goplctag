// Mozilla Public License 2.0, and LGPL License.

// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package goplctag

/*
#cgo pkg-config: libplctag
#include "libplctag.h"
#include "callbacks.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

func (x CreateExCallback) PassRef() (ref *C.go_plc_tag_create_ex_callback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if createExCallback6EE5B39CFunc == nil {
		createExCallback6EE5B39CFunc = x
	}
	return (*C.go_plc_tag_create_ex_callback)(C.go_plc_tag_create_ex_callback_6ee5b39c), nil
}

func (x CreateExCallback) PassValue() (ref C.go_plc_tag_create_ex_callback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if createExCallback6EE5B39CFunc == nil {
		createExCallback6EE5B39CFunc = x
	}
	return (C.go_plc_tag_create_ex_callback)(C.go_plc_tag_create_ex_callback_6ee5b39c), nil
}

func NewCreateExCallbackRef(ref unsafe.Pointer) *CreateExCallback {
	return (*CreateExCallback)(ref)
}

//export createExCallback6EE5B39C
func createExCallback6EE5B39C(ctagId C.int32_t, cevent C.int, cstatus C.int, cuserdata unsafe.Pointer) {
	if createExCallback6EE5B39CFunc != nil {
		tagId6ee5b39c := (int32)(ctagId)
		event6ee5b39c := (int32)(cevent)
		status6ee5b39c := (int32)(cstatus)
		userdata6ee5b39c := (unsafe.Pointer)(unsafe.Pointer(cuserdata))
		createExCallback6EE5B39CFunc(tagId6ee5b39c, event6ee5b39c, status6ee5b39c, userdata6ee5b39c)
		return
	}
	panic("callback func has not been set (race?)")
}

var createExCallback6EE5B39CFunc CreateExCallback

func (x RegisterExCallback) PassRef() (ref *C.go_plc_tag_register_ex_callback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if registerExCallback858E2960Func == nil {
		registerExCallback858E2960Func = x
	}
	return (*C.go_plc_tag_register_ex_callback)(C.go_plc_tag_register_ex_callback_858e2960), nil
}

func (x RegisterExCallback) PassValue() (ref C.go_plc_tag_register_ex_callback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if registerExCallback858E2960Func == nil {
		registerExCallback858E2960Func = x
	}
	return (C.go_plc_tag_register_ex_callback)(C.go_plc_tag_register_ex_callback_858e2960), nil
}

func NewRegisterExCallbackRef(ref unsafe.Pointer) *RegisterExCallback {
	return (*RegisterExCallback)(ref)
}

//export registerExCallback858E2960
func registerExCallback858E2960(ctagId C.int32_t, cevent C.int, cstatus C.int, cuserdata unsafe.Pointer) {
	if registerExCallback858E2960Func != nil {
		tagId858e2960 := (int32)(ctagId)
		event858e2960 := (int32)(cevent)
		status858e2960 := (int32)(cstatus)
		userdata858e2960 := (unsafe.Pointer)(unsafe.Pointer(cuserdata))
		registerExCallback858E2960Func(tagId858e2960, event858e2960, status858e2960, userdata858e2960)
		return
	}
	panic("callback func has not been set (race?)")
}

var registerExCallback858E2960Func RegisterExCallback

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}

func (x RegisterLoggerCallback) PassRef() (ref *C.go_plc_tag_register_logger_callback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if registerLoggerCallback63095D5AFunc == nil {
		registerLoggerCallback63095D5AFunc = x
	}
	return (*C.go_plc_tag_register_logger_callback)(C.go_plc_tag_register_logger_callback_63095d5a), nil
}

func (x RegisterLoggerCallback) PassValue() (ref C.go_plc_tag_register_logger_callback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if registerLoggerCallback63095D5AFunc == nil {
		registerLoggerCallback63095D5AFunc = x
	}
	return (C.go_plc_tag_register_logger_callback)(C.go_plc_tag_register_logger_callback_63095d5a), nil
}

func NewRegisterLoggerCallbackRef(ref unsafe.Pointer) *RegisterLoggerCallback {
	return (*RegisterLoggerCallback)(ref)
}

//export registerLoggerCallback63095D5A
func registerLoggerCallback63095D5A(ctagId C.int32_t, cdebugLevel C.int, cmessage *C.char) {
	if registerLoggerCallback63095D5AFunc != nil {
		tagId63095d5a := (int32)(ctagId)
		debugLevel63095d5a := (int32)(cdebugLevel)
		message63095d5a := packPCharString(cmessage)
		registerLoggerCallback63095D5AFunc(tagId63095d5a, debugLevel63095d5a, message63095d5a)
		return
	}
	panic("callback func has not been set (race?)")
}

var registerLoggerCallback63095D5AFunc RegisterLoggerCallback

// unpackPCharString copies the data from Go string as *C.char.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CString(str))
	allocs.Add(mem0)
	return (*C.char)(mem0), allocs
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// copyPUint8Bytes copies the data from Go slice as *C.uint8_t.
func copyPUint8Bytes(slice *sliceHeader) (*C.uint8_t, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfUint8Value) * slice.Len,
		Cap:  int(sizeOfUint8Value) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.uint8_t)(mem0), allocs
}

// allocUint8Memory allocates memory for type C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUint8Value))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfUint8Value = unsafe.Sizeof([1]C.uint8_t{})

// copyPCharBytes copies the data from Go slice as *C.char.
func copyPCharBytes(slice *sliceHeader) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfCharValue) * slice.Len,
		Cap:  int(sizeOfCharValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.char)(mem0), allocs
}

// allocCharMemory allocates memory for type C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfCharValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfCharValue = unsafe.Sizeof([1]C.char{})
