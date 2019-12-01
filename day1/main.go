package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		total += calcFuelOfFuel(n)
	}
	fmt.Println(total)
}

func calcFuel(mass int) int {
	var res float64
	res = math.Floor(float64(mass/3)) - 2
	total := int(res)
	return total
}

func calcFuelOfFuel(fuel int) int {
	temp := fuel
	result := 0
	for temp > 0 {
		v := calcFuel(temp)
		if v <= 0 {
			break
		}
		temp = v
		result += v
	}
	return result
}
