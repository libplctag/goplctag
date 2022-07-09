package goplctag

import (
	"fmt"
	"testing"
)



func TestCreateTag(t *testing.T){
	const (
		tagPath = "protocol=ab_eip&gateway=192.168.1.3&path=1,2&cpu=LGX&elem_size=1&elem_count=1&debug=1&name=BaseBOOL"
		timeout = 5000
	)
	tag := Create(tagPath, timeout)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}
}

func TestSetDebugLevel(t *testing.T){
	SetDebugLevel(DebugDetail)
}

func TestCheckLibVersion(t *testing.T){
	versionStatus := CheckLibVersion(2, 5, 1)
	if int32(versionStatus) != StatusOk {
		t.Errorf("Version does not match")
	}
}

func TestDintArray(t *testing.T){
	tagPath := "protocol=ab_eip&gateway=192.168.1.3&path=1,2&cpu=LGX&elem_count=128&name=BaseDINTArray"
	timeout := 5000

	SetDebugLevel(DebugNone)

	/* create the tag */
	tag := Create(tagPath, int32(timeout))
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", DecodeError(int32(tag)))
	}

	/* get the data */
	rc := Read(tag, int32(timeout))
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

	/* print out the data */
	for i := 0; i < int(elemCount); i++ {
		fmt.Printf("data[%v]=%v\n", i, GetInt32(tag, int32(i)*elemSize))
	}

	/* now test a write */
	for i := 0; i < int(elemCount); i++ {
		val := GetInt32(tag, int32(i)*elemSize)
		val++
		SetInt32(tag, int32(i)*elemSize, val)
	}
	rc = Write(tag, int32(timeout))
	if rc != StatusOk {
		Destroy(tag)
		t.Errorf("ERROR %s: Unable to write the data!\n", DecodeError(int32(tag)))		
	}

	/* get the data again*/
	rc = Read(tag, int32(timeout))
	if rc != StatusOk {
		Destroy(tag)
		t.Errorf("ERROR %s: Unable to read the data!\n", DecodeError(int32(tag)))		
	}

	/* print out the data */
	for i := 0; i < int(elemCount); i++ {
		fmt.Printf("data[%v]=%v\n", i, GetInt32(tag, int32(i)*elemSize))
	}

	Destroy(tag)

}