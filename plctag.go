/***************************************************************************
 *   Copyright (C) 2019 Aníbal Limón <limon.anibal@gmail.com>              *
 *                                                                         *
 * This software is available under either the Mozilla Public License      *
 * version 2.0 or the GNU LGPL version 2 (or later) license, whichever     *
 * you choose.                                                             *
 *                                                                         *
 * MPL 2.0:                                                                *
 *                                                                         *
 *   This Source Code Form is subject to the terms of the Mozilla Public   *
 *   License, v. 2.0. If a copy of the MPL was not distributed with this   *
 *   file, You can obtain one at http://mozilla.org/MPL/2.0/.              *
 *                                                                         *
 *                                                                         *
 * LGPL 2:                                                                 *
 *                                                                         *
 *   This program is free software; you can redistribute it and/or modify  *
 *   it under the terms of the GNU Library General Public License as       *
 *   published by the Free Software Foundation; either version 2 of the    *
 *   License, or (at your option) any later version.                       *
 *                                                                         *
 *   This program is distributed in the hope that it will be useful,       *
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of        *
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the         *
 *   GNU General Public License for more details.                          *
 *                                                                         *
 *   You should have received a copy of the GNU Library General Public     *
 *   License along with this program; if not, write to the                 *
 *   Free Software Foundation, Inc.,                                       *
 *   59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.             *
 ***************************************************************************/

package goplctag

/*
  #cgo pkg-config: libplctag
  #include <stdio.h>
  #include <stdlib.h>
  #include <libplctag.h>
*/
import "C"
import "unsafe"
  
const (
	
	StatusPending = 1
	StatusOk = 0
	ErrAbort = -1
	ErrBadConfig = -2
	ErrBadConnection = -3
	ErrBadData = -4
	ErrBadDevice = -5
	ErrBadGateway = -6
	ErrBadParam = -7
	ErrBadReply = -8
	ErrBadStatus = -9
	ErrClose = -10
	ErrCreate = -11
	ErrDuplicate = -12
	ErrEncode = -13
	ErrMutexDestroy = -14
	ErrMutexInit = -15
	ErrMutexLock = -16
	ErrMutexUnlock = -17
	ErrNotAllowed = -18
	ErrNotFound = -19
	ErrNotImplemented = -20
	ErrNoData = -21
	ErrNoMatch = -22
	ErrNoMem = -23
	ErrNoResources = -24
	ErrNullPtr = -25
	ErrOpen = -26
	ErrOutOfBounds = -27
	ErrRead = -28
	ErrRemoteErr = -29
	ErrThreadCreate = -30
	ErrThreadJoin = -31
	ErrTimeout = -32
	ErrTooLarge = -33
	ErrTooSmall = -34
	ErrUnsupported = -35
	ErrWinsock = -36
	ErrWrite = -37
	ErrPartial = -38
	ErrBusy = -39
	DebugNone = 0
	DebugError = 1
	DebugWarn = 2
	DebugInfo = 3
	DebugDetail = 4
	DebugSpew = 5
	EventReadStarted = 1
	EventReadCompleted = 2
	EventWriteStarted = 3
	EventWriteCompleted = 4
	EventAborted = 5
	EventDestroyed = 6
	EventCreated = 7
	EventMax = 8
)
  
  func DecodeError(err int) string {
	  cstr := C.plc_tag_decode_error(C.int(err))
	  return C.GoString(cstr)
  }
  
  func Create(attrib_str string, timeout int) int32 {
	  cattrib_str := C.CString(attrib_str)
	  result := C.plc_tag_create(cattrib_str, C.int(timeout))
	  C.free(unsafe.Pointer(cattrib_str))
	  return int32(result)
  }
  
  func Destroy(tag int32) int {
	  result := C.plc_tag_destroy(C.int32_t(tag))
	  return int(result)
  }
  
  func Lock(tag int32) int {
	  result := C.plc_tag_lock(C.int32_t(tag))
	  return int(result)
  }
  
  func Unlock(tag int32) int {
	  result := C.plc_tag_unlock(C.int32_t(tag))
	  return int(result)
  }
  
  func Abort(tag int32) int {
	  result := C.plc_tag_abort(C.int32_t(tag))
	  return int(result)
  }
  
  func Status(tag int32) int {
	  result := C.plc_tag_status(C.int32_t(tag))
	  return int(result)
  }
  
  func GetSize(tag int32) int {
	  result := C.plc_tag_get_size(C.int32_t(tag))
	  return int(result)
  }
  
  func Read(tag int32, timeout int) int {
	  result := C.plc_tag_read(C.int32_t(tag), C.int(timeout))
	  return int(result)
  }
  
  func Write(tag int32, timeout int) int {
	  result := C.plc_tag_write(C.int32_t(tag), C.int(timeout))
	  return int(result)
  }
  
  func GetUint64(tag int32, offset int) uint64 {
	  result := C.plc_tag_get_uint64(C.int32_t(tag), C.int(offset))
	  return uint64(result)
  }
  
  func SetUint64(tag int32, offset int, val uint64) int {
	  result := C.plc_tag_set_uint64(C.int32_t(tag), C.int(offset), C.uint64_t(val))
	  return int(result)
  }
  
  func GetInt64(tag int32, offset int) int64 {
	  result := C.plc_tag_get_int64(C.int32_t(tag), C.int(offset))
	  return int64(result)
  }
  
  func SetInt64(tag int32, offset int, val int64) int {
	  result := C.plc_tag_set_int64(C.int32_t(tag), C.int(offset), C.int64_t(val))
	  return int(result)
  }
  
  func GetUint32(tag int32, offset int) uint32 {
	  result := C.plc_tag_get_uint32(C.int32_t(tag), C.int(offset))
	  return uint32(result)
  }
  
  func SetUint32(tag int32, offset int, val uint32) int {
	  result := C.plc_tag_set_uint32(C.int32_t(tag), C.int(offset), C.uint32_t(val))
	  return int(result)
  }
  
  func GetInt32(tag int32, offset int) int32 {
	  result := C.plc_tag_get_int32(C.int32_t(tag), C.int(offset))
	  return int32(result)
  }
  
  func SetInt32(tag int32, offset int, val int32) int {
	  result := C.plc_tag_set_int32(C.int32_t(tag), C.int(offset), C.int32_t(val))
	  return int(result)
  }
  
  func GetUint16(tag int32, offset int) uint16 {
	  result := C.plc_tag_get_uint16(C.int32_t(tag), C.int(offset))
	  return uint16(result)
  }
  
  func SetUint16(tag int32, offset int, val uint16) int {
	  result := C.plc_tag_set_uint16(C.int32_t(tag), C.int(offset), C.uint16_t(val))
	  return int(result)
  }
  
  func GetInt16(tag int32, offset int) int16 {
	  result := C.plc_tag_get_int16(C.int32_t(tag), C.int(offset))
	  return int16(result)
  }
  
  func SetInt16(tag int32, offset int, val int16) int {
	  result := C.plc_tag_set_int16(C.int32_t(tag), C.int(offset), C.int16_t(val))
	  return int(result)
  }
  
  func GetUint8(tag int32, offset int) uint8 {
	  result := C.plc_tag_get_uint8(C.int32_t(tag), C.int(offset))
	  return uint8(result)
  }
  
  func SetUint8(tag int32, offset int, val uint8) int {
	  result := C.plc_tag_set_uint8(C.int32_t(tag), C.int(offset), C.uint8_t(val))
	  return int(result)
  }
  
  func GetInt8(tag int32, offset int) int8 {
	  result := C.plc_tag_get_int8(C.int32_t(tag), C.int(offset))
	  return int8(result)
  }
  
  func SetInt8(tag int32, offset int, val int8) int {
	  result := C.plc_tag_set_int8(C.int32_t(tag), C.int(offset), C.int8_t(val))
	  return int(result)
  }
  
  func GetFloat64(tag int32, offset int) float64 {
	  result := C.plc_tag_get_float64(C.int32_t(tag), C.int(offset))
	  return float64(result)
  }
  
  func SetFloat64(tag int32, offset int, val float64) int {
	  result := C.plc_tag_set_float64(C.int32_t(tag), C.int(offset), C.double(val))
	  return int(result)
  }
  
  func GetFloat32(tag int32, offset int) float32 {
	  result := C.plc_tag_get_float32(C.int32_t(tag), C.int(offset))
	  return float32(result)
  }
  
  func SetFloat32(tag int32, offset int, val float32) int {
	  result := C.plc_tag_set_float32(C.int32_t(tag), C.int(offset), C.float(val))
	  return int(result)
  }
  
  func GetBit(tag int32, offset int) int {
	  result := C.plc_tag_get_bit(C.int32_t(tag), C.int(offset))
	  return int(result)
  }
  
  func SetBit(tag int32, offset int, val int) int {
	  result := C.plc_tag_set_bit(C.int32_t(tag), C.int(offset), C.int(val))
	  return int(result)
  }
  