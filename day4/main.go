package main

import (
	"fmt"
	"strconv"
)

func main() {
	lower := 206938
	upper := 679128
	count := 0

	for i := lower; i <= upper; i++ {
		if valid(strconv.Itoa(i)) {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(valid("112233")) // true
	fmt.Println(valid("123444")) // false
	fmt.Println(valid("111122")) // true
	fmt.Println(valid("113444")) // true
}

func valid(input string) bool {
	adjacent := false
	for i := 0; i < len(input)-1; {
		if input[i] > input[i+1] {
			return false
		}
		if input[i] == input[i+1] {
			length := countNums(input, i)
			i += length - 1
			if length == 2 {
				adjacent = true
			}
		} else {
			i++
		}

	}

	return adjacent
}

func countNums(input string, idx int) int {
	count := 0
	c := input[idx]
	for idx < len(input) && c == input[idx] {
		idx++
		count++
	}
	return count
}
