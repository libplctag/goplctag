package goplctag

import (
	"testing"
)

func TestCreateTag(t *testing.T) {
	const (
		tagPath = protocol + gateway + plcType + "elem_size=1&elem_count=1&debug=1&name=BaseBOOL"
	)
	tag := Create(tagPath, TIMEOUT)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}
}

func TestSetDebugLevel(t *testing.T) {
	SetDebugLevel(DebugDetail)
}

func TestCheckLibVersion(t *testing.T) {
	versionStatus := CheckLibVersion(2, 5, 1)
	if int32(versionStatus) != StatusOk {
		t.Errorf("Version does not match")
	}
}

func TestDintArray(t *testing.T) {
	tagPath := protocol + gateway + plcType + "elem_count=128&name=BaseDINTArray"

	SetDebugLevel(DebugNone)

	/* create the tag */
	tag := Create(tagPath, int32(TIMEOUT))
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}

	/* get the data */
	rc := Read(tag, int32(TIMEOUT))
	if rc != StatusOk {
		Destroy(tag)
		t.Errorf("ERROR %s: Unable to read the data!\n", DecodeError(int32(tag)))
	}

	/* get the tag size and element size.
	Do this _AFTER_ reading the tag otherwise we may not know how big the tag is! */
	elemSize := GetIntAttribute(tag, "elem_size", 0)
	elemCount := GetIntAttribute(tag, "elem_count", 0)

	//fmt.Printf("Count %v, %v bytes\n", elemCount, elemSize)

	if elemCount != 128 {
		t.Errorf("Wrong element count, expecting 128 elements!\n")
	}

	if elemSize != 4 {
		t.Errorf("Wrong element size, expecting 4 bytes!\n")
	}

	arraySize := GetSize(tag)

	if arraySize != elemCount*elemSize {
		t.Errorf("Wrong array size, expecting 512 bytes!\n")
	}

	/* print out the data */
	//for i := 0; i < int(elemCount); i++ {
	//	fmt.Printf("data[%v]=%v\n", i, GetInt32(tag, int32(i)*elemSize))
	//}

	/* now test write */
	for i := 0; i < int(elemCount); i++ {
		val := GetInt32(tag, int32(i)*elemSize)
		val++
		SetInt32(tag, int32(i)*elemSize, val)
	}
	rc = Write(tag, int32(TIMEOUT))
	if rc != StatusOk {
		Destroy(tag)
		t.Errorf("ERROR %s: Unable to write the data!\n", DecodeError(int32(tag)))
	}

	/* get the data again*/
	rc = Read(tag, int32(TIMEOUT))
	if rc != StatusOk {
		Destroy(tag)
		t.Errorf("ERROR %s: Unable to read the data!\n", DecodeError(int32(tag)))
	}

	/* print out the data */
	//for i := 0; i < int(elemCount); i++ {
	//	fmt.Printf("data[%v]=%v\n", i, GetInt32(tag, int32(i)*elemSize))
	//}

	Destroy(tag)
}

func TestBoolTag(t *testing.T) {
	const (
		tagPath = protocol + gateway + plcType + "elem_size=1&elem_count=1&debug=1&name=BaseBOOL"
	)
	/* create the tag */
	tag := Create(tagPath, TIMEOUT)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}

	/* everything OK? */
	if rc := Status(tag); rc != StatusOk {
		t.Errorf("ERROR %s: Error setting up tag internal state.\n", DecodeError(rc))
	}

	/* get the data */
	if rc := Read(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to read the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	b := GetUint8(tag, 0)
	if b > 0 {
		b = 0
	} else {
		b = 255
	}

	SetUint8(tag, 0, b)
	if rc := Write(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to write the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	bit := GetBit(tag, 0)
	if bit > 0 {
		bit = 0
	} else {
		bit = 1
	}

	SetBit(tag, 0, bit)
	if rc := Write(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to write the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	bit = GetBit(tag, 0)
	println("BaseBOOL: ", bit)
}

func TestRealTag(t *testing.T) {
	const (
		tagPath = protocol + gateway + plcType + "elem_size=1&elem_count=1&debug=1&name=BaseREAL"
	)
	/* create the tag */
	tag := Create(tagPath, TIMEOUT)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}

	/* everything OK? */
	if rc := Status(tag); rc != StatusOk {
		t.Errorf("ERROR %s: Error setting up tag internal state.\n", DecodeError(rc))
	}

	/* get the data */
	if rc := Read(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to read the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 := GetFloat32(tag, 0)

	if t1 > 2.5 {
		t1 = 1.5
	} else {
		t1 = 3.5
	}

	SetFloat32(tag, 0, t1)
	if rc := Write(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to write the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 = GetFloat32(tag, 0)
	println("BaseREAL: ", t1)
}

func TestSintTag(t *testing.T) {
	const (
		tagPath = protocol + gateway + plcType + "elem_size=1&elem_count=1&debug=1&name=BaseSINT"
	)
	/* create the tag */
	tag := Create(tagPath, TIMEOUT)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}

	/* everything OK? */
	if rc := Status(tag); rc != StatusOk {
		t.Errorf("ERROR %s: Error setting up tag internal state.\n", DecodeError(rc))
	}

	/* get the data */
	if rc := Read(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to read the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 := GetInt8(tag, 0) // BUG needs fixing on auto-generation, manually fixed for now

	if t1 > 1 {
		t1 = -55
	} else {
		t1 = 55
	}

	SetInt8(tag, 0, t1)
	if rc := Write(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to write the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 = GetInt8(tag, 0)
	println("BaseSINT: ", t1)
}

func TestIntTag(t *testing.T) {
	const (
		tagPath = protocol + gateway + plcType + "elem_size=1&elem_count=1&debug=1&name=BaseINT"
	)
	/* create the tag */
	tag := Create(tagPath, TIMEOUT)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}

	/* everything OK? */
	if rc := Status(tag); rc != StatusOk {
		t.Errorf("ERROR %s: Error setting up tag internal state.\n", DecodeError(rc))
	}

	/* get the data */
	if rc := Read(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to read the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 := GetInt16(tag, 0)

	if t1 > 1 {
		t1 = -6512
	} else {
		t1 = 6512
	}

	SetInt16(tag, 0, t1)
	if rc := Write(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to write the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 = GetInt16(tag, 0)
	println("BaseINT: ", t1)
}

func TestDintTag(t *testing.T) {
	const (
		tagPath = protocol + gateway + plcType + "elem_size=1&elem_count=1&debug=1&name=BaseDINT"
	)
	/* create the tag */
	tag := Create(tagPath, TIMEOUT)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}

	/* everything OK? */
	if rc := Status(tag); rc != StatusOk {
		t.Errorf("ERROR %s: Error setting up tag internal state.\n", DecodeError(rc))
	}

	/* get the data */
	if rc := Read(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to read the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 := GetInt32(tag, 0)

	if t1 > 1 {
		t1 = -2147483648
	} else {
		t1 = 2147483647
	}

	SetInt32(tag, 0, t1)
	if rc := Write(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to write the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 = GetInt32(tag, 0)
	println("BaseDINT: ", t1)
}

func TestLintTag(t *testing.T) {
	const (
		tagPath = protocol + gateway + plcType + "elem_size=1&elem_count=1&debug=1&name=BaseLINT"
	)
	/* create the tag */
	tag := Create(tagPath, TIMEOUT)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}

	/* everything OK? */
	if rc := Status(tag); rc != StatusOk {
		t.Errorf("ERROR %s: Error setting up tag internal state.\n", DecodeError(rc))
	}

	/* get the data */
	if rc := Read(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to read the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 := GetInt64(tag, 0)

	if t1 > 1 {
		t1 = -566906435
	} else {
		t1 = 566906435
	}

	SetInt64(tag, 0, t1)
	if rc := Write(tag, TIMEOUT); rc != StatusOk {
		t.Errorf("ERROR: Unable to write the data! Got error code %d: %s\n", rc, DecodeError(rc))
	}

	t1 = GetInt64(tag, 0)
	println("BaseLINT: ", t1)
}
