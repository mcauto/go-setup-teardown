package unittest

import "testing"

type setup func() error

type teardown func() error

// Setup 테스트 setup, teardown
func Setup(s setup, td teardown, t *testing.T) func(t *testing.T) {
	s()
	// return teardown
	return func(t *testing.T) {
		td()
	}
}

