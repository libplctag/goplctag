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
import "unsafe"

// CreateExCallback type as declared in goplctag/callbacks.h:32
type CreateExCallback func(tagId int32, event int32, status int32, userdata unsafe.Pointer)

// RegisterExCallback type as declared in goplctag/callbacks.h:36
type RegisterExCallback func(tagId int32, event int32, status int32, userdata unsafe.Pointer)

// RegisterLoggerCallback type as declared in goplctag/callbacks.h:40
type RegisterLoggerCallback func(tagId int32, debugLevel int32, message string)
