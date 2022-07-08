package goplctag

import (
	plctag "goplctag"
	"testing"
)

const (
	TAG_PATH     = "protocol=ab_eip&gateway=192.168.1.3&path=1,2&cpu=LGX&elem_size=1&elem_count=1&debug=1&name=BaseBOOLL"
	DATA_TIMEOUT = 5000
)

func TestCreateTag(t *testing.T){
	tag := plctag.Create(TAG_PATH, DATA_TIMEOUT)
	if tag < 0 {
		t.Errorf("ERROR %s: Could not create tag!\n", plctag.DecodeError(int32(tag)))
	}	 
}