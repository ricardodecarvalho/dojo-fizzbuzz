package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, test := range tests {
		t.Run("Test FizzBuzz: "+strconv.Itoa(test.input), func(t *testing.T) {
			res := FizzBuzz(test.input)
			if !reflect.DeepEqual(res, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, res)
			}
		})
	}
}
