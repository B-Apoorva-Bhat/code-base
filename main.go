package main

import (
	"fmt"
	"reflect"
	"strings"
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
func add(a int) int {
	for a/10 > 0 {
		a = a/10 + a%10
	}
	// fmt.Println(a)
	return a
}
func convertor(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
func main() {
	// r := 0
	// s := 0

	var number int
	fmt.Print("enter number\n")
	_, err := fmt.Scanln(&number)

	// x := "015"
	// i := strings.Index(x, "0")

	// fmt.Println("Index: ", i)
	// if i == 0 {
	// 	fmt.Println("error")
	// }
	// for number != 0 {
	// 	r = number % 10
	// 	s += r
	// 	number = number / 10
	// }

	// if s > 9 {
	// 	rem := 0
	// 	result := 0
	// 	for s != 0 {
	// 		rem = s % 10
	// 		result += rem
	// 		s = s / 10

	// 	}
	// 	fmt.Println(result)
	// }
	if err != nil {
		fmt.Println("Enter a valid number!!", err)
		return
	}

	fmt.Printf("sum is %d \n", add(number))

	var digits []int

	for number > 0 {
		digits = append(digits, number%10)

		number = number / 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	fmt.Println("Number into digits : ", digits)
	stringvalue := convertor(digits, "")
	fmt.Println("Integer to string : ", stringvalue)
	fmt.Println("t", reflect.TypeOf(stringvalue))
	fmt.Printf("Type is %T", stringvalue)

}
