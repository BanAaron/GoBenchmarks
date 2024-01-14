package main

import "fmt"

// Hello returns a "Hello, World!" with the provided name
//
// if no name is provided "World" will be used
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
