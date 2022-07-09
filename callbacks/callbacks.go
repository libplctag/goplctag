// Mozilla Public License 2.0, and LGPL License.

// WARNING: This file has automatically been generated on Sat, 09 Jul 2022 16:25:18 CDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package callbacks

/*
#cgo pkg-config: libplctag
#include "libplctag.h"
#include "callbacks.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// CreateEx function as declared in callbacks/callbacks.h:33
func CreateEx(attribStr string, callbackFunction CreateExCallback, userdata unsafe.Pointer, timeout int32) int32 {
	cattribStr, cattribStrAllocMap := unpackPCharString(attribStr)
	ccallbackFunction, ccallbackFunctionAllocMap := callbackFunction.PassValue()
	cuserdata, cuserdataAllocMap := userdata, cgoAllocsUnknown
	ctimeout, ctimeoutAllocMap := (C.int)(timeout), cgoAllocsUnknown
	__ret := C.go_plc_tag_create_ex(cattribStr, ccallbackFunction, cuserdata, ctimeout)
	runtime.KeepAlive(ctimeoutAllocMap)
	runtime.KeepAlive(cuserdataAllocMap)
	runtime.KeepAlive(ccallbackFunctionAllocMap)
	runtime.KeepAlive(cattribStrAllocMap)
	__v := (int32)(__ret)
	return __v
}

// RegisterCallbackEx function as declared in callbacks/callbacks.h:37
func RegisterCallbackEx(tagId int32, callbackFunction RegisterExCallback, userdata unsafe.Pointer) int32 {
	ctagId, ctagIdAllocMap := (C.int32_t)(tagId), cgoAllocsUnknown
	ccallbackFunction, ccallbackFunctionAllocMap := callbackFunction.PassValue()
	cuserdata, cuserdataAllocMap := userdata, cgoAllocsUnknown
	__ret := C.go_plc_tag_register_callback_ex(ctagId, ccallbackFunction, cuserdata)
	runtime.KeepAlive(cuserdataAllocMap)
	runtime.KeepAlive(ccallbackFunctionAllocMap)
	runtime.KeepAlive(ctagIdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// RegisterLogger function as declared in callbacks/callbacks.h:41
func RegisterLogger(callbackFunction RegisterLoggerCallback) int32 {
	ccallbackFunction, ccallbackFunctionAllocMap := callbackFunction.PassValue()
	__ret := C.go_plc_tag_register_logger(ccallbackFunction)
	runtime.KeepAlive(ccallbackFunctionAllocMap)
	__v := (int32)(__ret)
	return __v
}