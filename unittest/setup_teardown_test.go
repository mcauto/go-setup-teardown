package unittest_test

import (
	"fmt"
	"portd/unittest"
	"testing"
)

func setup() error {
	fmt.Println("test")
	return nil
}

func teardown() error {
	fmt.Println("teardown")
	return nil
}

func TestSetup(t *testing.T) {
	teardown := unittest.Setup(setup, teardown, t)
	defer teardown(t)
}

