package main

import "testing"

func AssertEqual(t *testing.T, a any, b any) {
	if a != b {
		t.Errorf("%v and %v are not equal", a, b)
	}
}
