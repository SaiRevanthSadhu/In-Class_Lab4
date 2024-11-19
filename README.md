# In Class Lab4 Assignment

### Below is the documentation for the each function in the util package

### IsPrime checks whether a given integer is a prime number.

### A prime number is greater than 1 and divisible only by 1 and itself.

### Returns true if the number is prime, otherwise false.

### Edge cases include handling negative numbers and non-integers.

## func IsPrime(n int) bool

### ReverseString reverses the characters in a given string.

### For example, "hello" becomes "olleh".

### Handles edge cases like empty strings and single-character strings.

## func ReverseString(s string) string

### SumOfSlice calculates the sum of all integers in a given slice.

### Returns 0 for an empty slice.

### Handles negative numbers in the slice and large slices efficiently.

## func SumOfSlice(nums []int) int

### Factorial calculates the factorial of a given non-negative integer n.

### The factorial of n (n!) is the product of all positive integers ≤ n.

### Returns 1 for n = 0, and -1 for negative inputs as an invalid case.

## func Factorial(n int) int

## 1. Challenges Encountered

### Identifying Edge Cases:

#### It was tricky to determine all edge cases for functions like IsPrime (e.g., handling 0, 1, and negative numbers).

#### For SumOfSlice, testing large slices and empty slices required special attention to performance and correctness.

### Handling Invalid Inputs:

#### Deciding how functions like Factorial should behave with negative inputs required careful thought (e.g., returning -1 for invalid inputs).

### Benchmark Setup:

#### For benchmark tests, creating input datasets of varying sizes and ensuring consistency across runs took time.

## 2. Table-Driven Tests and Subtests

### How They Helped:

#### Table-driven tests ensured systematic coverage of multiple input cases for each function.

#### Subtests with t.Run() provided clear structure and allowed for better debugging. Each scenario could be tested and debugged independently without interference from other cases.

### Benefits:

#### Made it easier to spot failures for specific inputs.

#### Simplified test maintenance as adding new cases required only modifying the test table.

## 3. Key Findings from Benchmark Tests

### IsPrime:

#### Performance scaled poorly with very large numbers, especially those requiring numerous divisibility checks.

#### Optimization with early exits (e.g., skipping even numbers after 2) significantly improved efficiency.

### ReverseString:

#### Performed well for short strings, but performance degraded linearly with string length.

#### Memory usage was efficient due to in-place character reversal.

### SumOfSlice:

#### Handled large slices effectively. Performance was linear with respect to the size of the input slice, as expected.

### Factorial:

#### Performance was quick for small values but highlighted potential overflow issues with very large inputs due to Go’s int type limitations.

## 4. Importance of Test Coverage

### Test Coverage Achieved:

#### Achieved over 90% coverage by addressing edge cases and invalid inputs.

#### Uncovered areas needing improvement, such as handling unusual inputs or extreme values.

### Reflection on Importance:

#### High test coverage ensures reliability, as every critical code path is verified.

#### Writing tests revealed hidden bugs (e.g., incorrect handling of edge cases in IsPrime and Factorial).

#### Helps build confidence in code quality during refactoring or when adding new features.

### Areas for Improvement:

#### Improve coverage for unexpected inputs (e.g., very large numbers, empty strings).

#### Consider adding stress tests to validate performance under extreme conditions.

#### Enhance error handling to return more informative feedback instead of generic errors like -1.