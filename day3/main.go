package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	row int
	col int
}

// for ur own sanity do not read this code =]
func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	paths := make([]string, 2)
	idx := 0
	for s.Scan() {
		var input string
		_, err := fmt.Sscanf(s.Text(), "%s", &input)
		if err != nil {
			log.Fatal(err)
		}
		paths[idx] = input
		idx++
	}
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	generateIntervals(m1, paths[0])
	generateIntervals(m2, paths[1])
	minSteps := float64(100000000)
	// distance := float64(1000000000)
	// for k := range m1 {
	// 	if m2[k] != 0 {
	// 		points := strings.Split(k, "_")
	// 		x, _ := strconv.Atoi(points[0][3:])
	// 		y, _ := strconv.Atoi(points[1][3:])
	// 		if x == 0 && y == 0 {
	// 			continue
	// 		} else {
	// 			result := (math.Abs(float64(0-x)) + math.Abs(float64(0-y)))
	// 			distance = math.Min(float64(distance), float64(result))
	// 		}
	// 	}
	// }
	// fmt.Println(distance)
	for k := range m1 {
		if m2[k] != 0 {
			// points := strings.Split(k, "_")
			//fmt.Println("map1 (", k, ") = ", m1[k], " || map2 = ", m2[k])
			minSteps = math.Min(minSteps, float64(m1[k]+m2[k]))
		}
	}

	fmt.Println(minSteps)
}

func generateIntervals(m map[string]int, moves string) {
	tokens := strings.Split(moves, ",")
	// row2_col3
	lastPos := pair{0, 0}
	currStep := 0
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		direction := token[0:1]
		distance, _ := strconv.Atoi(token[1:])

		var currPos pair
		if direction == "R" {
			targetCol := lastPos.col + distance
			currCol := lastPos.col

			for currCol <= targetCol {
				p := fmt.Sprintf("row%s_col%s", strconv.Itoa(lastPos.row), strconv.Itoa(currCol))
				if !(currCol == 0 && lastPos.row == 0) && m[p] == 0 {
					m[p] = currStep
				}
				if currCol != targetCol-1 {
					currStep++
				}
				currCol++
			}
			currPos = pair{lastPos.row, lastPos.col + distance}
		} else if direction == "L" {
			targetCol := lastPos.col - distance
			currCol := lastPos.col

			for currCol >= targetCol {
				p := fmt.Sprintf("row%s_col%s", strconv.Itoa(lastPos.row), strconv.Itoa(currCol))

				if !(currCol == 0 && lastPos.row == 0) && m[p] == 0 {
					m[p] = currStep
				}
				if currCol != targetCol+1 {
					currStep++
				}
				currCol--
			}
			currPos = pair{lastPos.row, lastPos.col - distance}
		} else if direction == "U" {
			targetRow := lastPos.row + distance
			currRow := lastPos.row

			for currRow <= targetRow {
				p := fmt.Sprintf("row%s_col%s", strconv.Itoa(currRow), strconv.Itoa(lastPos.col))
				if !(lastPos.col == 0 && currRow == 0) && m[p] == 0 {
					m[p] = currStep
				}
				if currRow != targetRow-1 {
					currStep++
				}
				currRow++
			}
			currPos = pair{lastPos.row + distance, lastPos.col}
		} else if direction == "D" {
			targetRow := lastPos.row - distance
			currRow := lastPos.row

			for currRow >= targetRow {
				p := fmt.Sprintf("row%s_col%s", strconv.Itoa(currRow), strconv.Itoa(lastPos.col))

				if !(lastPos.col == 0 && currRow == 0) && m[p] == 0 {
					m[p] = currStep
				}
				if currRow != targetRow+1 {
					currStep++
				}
				currRow--
			}
			currPos = pair{lastPos.row - distance, lastPos.col}
		}
		lastPos = currPos
	}
}
