package unittest_test

import (
	"fmt"
	"portd/unittest"
	"testing"
)

func (ut *unittestTest) setup() error {
	fmt.Println("test")
	return nil
}

func (ut *unittestTest) teardown() error {
	fmt.Println("teardown")
	return nil
}

func TestSetup(t *testing.T) {
	teardown := unittest.Setup(ut.setup, ut.teardown, t)
	defer teardown(t)
}
