package util

// IsPrime checks if a given integer is a prime number.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// ReverseString reverses a given string.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// SumOfSlice returns the sum of all integers in a slice.
func SumOfSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
