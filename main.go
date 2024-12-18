package main

import (
	"fmt"

	"github.com/SaiRevanthSadhu/csdd1008_fall24/Assignment4/util/util"
)

func main() {
	fmt.Println("Is 7 prime?", util.IsPrime(7))
	fmt.Println("Reverse of 'hello':", util.ReverseString("hello"))
	fmt.Println("Sum of slice [1, 2, 3, 4]:", util.SumOfSlice([]int{1, 2, 3, 4}))
	fmt.Println("Factorial of 5:", util.Factorial(5))
}
