package main

import (
	"fmt"
)

// func sum(number int) int {

// 	r := 0
// 	s := 0
// 	for number != 0 {
// 		r = number % 10
// 		s += r
// 		number = number / 10
// 	}

// 	if s > 9 {
// 		return sum(s)
// 	}
// 	return s
// }
func main() {
	r := 0
	s := 0
	var number int
	fmt.Print("enter number")
	fmt.Scanln(&number)
	for number != 0 {
		r = number % 10
		s += r
		number = number / 10
	}

	if s > 9 {
		rem := 0
		result := 0
		for s != 0 {
			rem = s % 10
			result += rem
			s = s / 10

		}
		fmt.Println(result)
	}
	// fmt.Printf("sum is %d", sum(number))
}
