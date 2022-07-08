// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Fri, 08 Jul 2022 14:56:03 CDT.
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
import "unsafe"

// CreateExCallback type as declared in goplctag/callbacks.h:32
type CreateExCallback func(tagId int32, event int32, status int32, userdata unsafe.Pointer)

// RegisterExCallback type as declared in goplctag/callbacks.h:36
type RegisterExCallback func(tagId int32, event int32, status int32, userdata unsafe.Pointer)

// RegisterLoggerCallback type as declared in goplctag/callbacks.h:40
type RegisterLoggerCallback func(tagId int32, debugLevel int32, message string)
