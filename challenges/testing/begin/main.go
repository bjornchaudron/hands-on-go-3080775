// challenges/testing/begin/main.go
package main

import "unicode"

type letterCounter struct{ identifier string }

func (l letterCounter) name() string {
	return l.identifier
}

func (l letterCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if unicode.IsLetter(char) {
			result++
		}
	}
	return result
}

type numberCounter struct{ designation string }

func (n numberCounter) name() string {
	return n.designation
}

func (n numberCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if unicode.IsNumber(char) {
			result++
		}
	}
	return result
}

type symbolCounter struct{ label string }

func (s symbolCounter) name() string {
	return s.label
}

func (s symbolCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			result++
		}
	}
	return result
}
