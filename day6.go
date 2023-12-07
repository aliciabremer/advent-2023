package main

import (
    "fmt"
	// "bufio"
	// "os"
	// "strings"
	// "strconv"
	"math"
)


func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func part1(times []int, dist []int) int {

	prod := 1
	for ind, t := range times {
		ways := 0
		for i := 1; i < 59; i++ {
			calc := (t - i) * i
			if calc >= dist[ind] {
				ways++
			}
		}
		prod *= ways
	}
	return prod
}

func part2(times int, dist int) int {

	quad := times * times - 4 * dist

	sqrt := math.Sqrt(float64(quad))

	minus := int((float64(times) - sqrt)/2+1)

	plus := int((float64(times) + sqrt)/2)

	// plus := times + sqrt

	fmt.Println((times - minus) * minus)
	fmt.Println((times - plus) * plus)

	return plus-minus+1
}

func main() {

	times := []int{59,68, 82, 74}

	dist := []int{543, 1020, 1664, 1022}

    sum := part1(times,dist)
	fmt.Println(sum)

	time := 59688274

	dist2 := 543102016641022

	su2m := part2(time, dist2)
	fmt.Println(su2m)
}