// Mozilla Public License 2.0, and LGPL License.

// WARNING: This file has automatically been generated on Sat, 09 Jul 2022 13:09:20 CDT.
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

const (
	// StatusPending as defined in include/libplctag.h:66
	StatusPending = 1
	// StatusOk as defined in include/libplctag.h:67
	StatusOk = 0
	// ErrAbort as defined in include/libplctag.h:69
	ErrAbort = -1
	// ErrBadConfig as defined in include/libplctag.h:70
	ErrBadConfig = -2
	// ErrBadConnection as defined in include/libplctag.h:71
	ErrBadConnection = -3
	// ErrBadData as defined in include/libplctag.h:72
	ErrBadData = -4
	// ErrBadDevice as defined in include/libplctag.h:73
	ErrBadDevice = -5
	// ErrBadGateway as defined in include/libplctag.h:74
	ErrBadGateway = -6
	// ErrBadParam as defined in include/libplctag.h:75
	ErrBadParam = -7
	// ErrBadReply as defined in include/libplctag.h:76
	ErrBadReply = -8
	// ErrBadStatus as defined in include/libplctag.h:77
	ErrBadStatus = -9
	// ErrClose as defined in include/libplctag.h:78
	ErrClose = -10
	// ErrCreate as defined in include/libplctag.h:79
	ErrCreate = -11
	// ErrDuplicate as defined in include/libplctag.h:80
	ErrDuplicate = -12
	// ErrEncode as defined in include/libplctag.h:81
	ErrEncode = -13
	// ErrMutexDestroy as defined in include/libplctag.h:82
	ErrMutexDestroy = -14
	// ErrMutexInit as defined in include/libplctag.h:83
	ErrMutexInit = -15
	// ErrMutexLock as defined in include/libplctag.h:84
	ErrMutexLock = -16
	// ErrMutexUnlock as defined in include/libplctag.h:85
	ErrMutexUnlock = -17
	// ErrNotAllowed as defined in include/libplctag.h:86
	ErrNotAllowed = -18
	// ErrNotFound as defined in include/libplctag.h:87
	ErrNotFound = -19
	// ErrNotImplemented as defined in include/libplctag.h:88
	ErrNotImplemented = -20
	// ErrNoData as defined in include/libplctag.h:89
	ErrNoData = -21
	// ErrNoMatch as defined in include/libplctag.h:90
	ErrNoMatch = -22
	// ErrNoMem as defined in include/libplctag.h:91
	ErrNoMem = -23
	// ErrNoResources as defined in include/libplctag.h:92
	ErrNoResources = -24
	// ErrNullPtr as defined in include/libplctag.h:93
	ErrNullPtr = -25
	// ErrOpen as defined in include/libplctag.h:94
	ErrOpen = -26
	// ErrOutOfBounds as defined in include/libplctag.h:95
	ErrOutOfBounds = -27
	// ErrRead as defined in include/libplctag.h:96
	ErrRead = -28
	// ErrRemoteErr as defined in include/libplctag.h:97
	ErrRemoteErr = -29
	// ErrThreadCreate as defined in include/libplctag.h:98
	ErrThreadCreate = -30
	// ErrThreadJoin as defined in include/libplctag.h:99
	ErrThreadJoin = -31
	// ErrTimeout as defined in include/libplctag.h:100
	ErrTimeout = -32
	// ErrTooLarge as defined in include/libplctag.h:101
	ErrTooLarge = -33
	// ErrTooSmall as defined in include/libplctag.h:102
	ErrTooSmall = -34
	// ErrUnsupported as defined in include/libplctag.h:103
	ErrUnsupported = -35
	// ErrWinsock as defined in include/libplctag.h:104
	ErrWinsock = -36
	// ErrWrite as defined in include/libplctag.h:105
	ErrWrite = -37
	// ErrPartial as defined in include/libplctag.h:106
	ErrPartial = -38
	// ErrBusy as defined in include/libplctag.h:107
	ErrBusy = -39
	// DebugNone as defined in include/libplctag.h:130
	DebugNone = 0
	// DebugError as defined in include/libplctag.h:131
	DebugError = 1
	// DebugWarn as defined in include/libplctag.h:132
	DebugWarn = 2
	// DebugInfo as defined in include/libplctag.h:133
	DebugInfo = 3
	// DebugDetail as defined in include/libplctag.h:134
	DebugDetail = 4
	// DebugSpew as defined in include/libplctag.h:135
	DebugSpew = 5
	// EventReadStarted as defined in include/libplctag.h:269
	EventReadStarted = 1
	// EventReadCompleted as defined in include/libplctag.h:270
	EventReadCompleted = 2
	// EventWriteStarted as defined in include/libplctag.h:272
	EventWriteStarted = 3
	// EventWriteCompleted as defined in include/libplctag.h:273
	EventWriteCompleted = 4
	// EventAborted as defined in include/libplctag.h:275
	EventAborted = 5
	// EventDestroyed as defined in include/libplctag.h:277
	EventDestroyed = 6
	// EventCreated as defined in include/libplctag.h:279
	EventCreated = 7
	// EventMax as defined in include/libplctag.h:281
	EventMax = 8
)
