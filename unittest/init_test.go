package unittest_test

import "testing"

type unittestTest struct {
}

var (
	ut *unittestTest
)

// can not override setup & teardown
func init() {
	ut = &unittestTest{}
}
