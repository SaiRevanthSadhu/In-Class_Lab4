package util

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Negative number", -5, false},
		{"Zero", 0, false},
		{"One", 1, false},
		{"Prime number", 7, true},
		{"Non-prime number", 10, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsPrime(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(7919) // A large prime number for testing.
	}
}
