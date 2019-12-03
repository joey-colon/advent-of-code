package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var result string
	for s.Scan() {
		var input string
		_, err := fmt.Sscanf(s.Text(), "%s", &input)
		if err != nil {
			log.Fatal(err)
		}
		result = calculateWrapper(input)
	}
	fmt.Println(result)
}

func calculateWrapper(input string) string {
	arr := strings.Split(input, ",")
	target := 19690720
	arr[1] = "12"
	arr[2] = "2"

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			currArr := make([]string, len(arr))
			copy(currArr, arr)
			currArr[1] = strconv.Itoa(noun)
			currArr[2] = strconv.Itoa(verb)
			if calculate(currArr, target) {
				return strings.Join(currArr, ",")
			}
		}
	}
	return ""
}

func calculate(arr []string, t int) bool {
	ptr := 0
	for ptr < len(arr) {
		if arr[ptr] == "1" {
			add(arr, ptr)
			ptr += 4
		} else if arr[ptr] == "2" {
			multiply(arr, ptr)
			ptr += 4
		} else if arr[ptr] == "99" {
			break
		} else {
			ptr++
		}
	}
	return arr[0] == strconv.Itoa(t)
}

// bleh
func add(arr []string, ptr int) {
	// ptr+1 and ptr+2 will contain the idx's we should perform computation
	// ptr+3 will contain the idx where we want to store the result.
	firstIdxString, secondIdxString, placementIdxString := arr[ptr+1], arr[ptr+2], arr[ptr+3]
	firstIdx, _ := strconv.Atoi(firstIdxString)
	secondIdx, _ := strconv.Atoi(secondIdxString)
	placementIdx, _ := strconv.Atoi(placementIdxString)
	val1, _ := strconv.Atoi(arr[firstIdx])
	val2, _ := strconv.Atoi(arr[secondIdx])

	arr[placementIdx] = strconv.Itoa(val1 + val2)
}

func multiply(arr []string, ptr int) {
	// ptr+1 and ptr+2 will contain the idx's we should perform computation
	// ptr+3 will contain the idx where we want to store the result.
	firstIdxString, secondIdxString, placementIdxString := arr[ptr+1], arr[ptr+2], arr[ptr+3]
	firstIdx, _ := strconv.Atoi(firstIdxString)
	secondIdx, _ := strconv.Atoi(secondIdxString)
	placementIdx, _ := strconv.Atoi(placementIdxString)
	val1, _ := strconv.Atoi(arr[firstIdx])
	val2, _ := strconv.Atoi(arr[secondIdx])

	arr[placementIdx] = strconv.Itoa(val1 * val2)
}
