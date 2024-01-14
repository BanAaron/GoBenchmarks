package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("aaron", func(t *testing.T) {
		res := Hello("Aaron")
		target := "Hello, Aaron!"
		if res != target {
			t.Errorf("%s is not equal to %s", res, target)
		}
	})
	t.Run("empty string", func(t *testing.T) {
		res := Hello("")
		target := "Hello, World!"
		if res != target {
			t.Errorf("%s is not equal to %s", res, target)
		}
	})
	t.Run("chinese", func(t *testing.T) {
		res := Hello("亚纶")
		target := "Hello, 亚纶!"
		if res != target {
			t.Errorf("%s is not equal to %s", res, target)
		}
	})
}
