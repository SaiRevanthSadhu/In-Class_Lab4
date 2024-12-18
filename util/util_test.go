package util

import (
	"strconv"
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

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
		{"Palindrome", "madam", "madam"},
		{"Regular string", "hello", "olleh"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReverseString(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
func TestSumOfSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{5}, 5},
		{"Multiple elements", []int{1, 2, 3, 4}, 10},
		{"Negative numbers", []int{-1, -2, -3}, -6},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SumOfSlice(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Negative number", -3, -1},
		{"Zero", 0, 1},
		{"One", 1, 1},
		{"Small number", 5, 120},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Factorial(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// BenchmarkIsPrime benchmarks the IsPrime function for multiple primes.
func BenchmarkIsPrime(b *testing.B) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	for _, prime := range primes {
		// Use strconv.Itoa to convert the integer to string for naming the benchmark
		b.Run("Prime_"+strconv.Itoa(prime), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsPrime(prime)
			}
		})
	}
}

func BenchmarkReverseString(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < b.N; i++ {
		ReverseString(str)
	}
}

func BenchmarkSumOfSlice(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	for i := 0; i < b.N; i++ {
		SumOfSlice(nums)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}
