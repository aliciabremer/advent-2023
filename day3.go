package main

import (
    "fmt"
	"bufio"
	"os"
	// "strings"
	// "strconv"
)

type Coord struct {
	x int
	y int
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func part1(scanner *bufio.Scanner) int {
	sum := 0

	my_map := make(map[int]([][]Coord))
	symbols := make([]Coord, 0)

	y := 0
	for scanner.Scan() {
		s := scanner.Text()

		len := len(s)
		
		x := 0

		checking_num := true
		num := 0

		lst := make([]Coord, 0)

		for x < len {
			if s[x] >= '0' && s[x] <= '9' {
				if !checking_num {
					num = 0
					lst = make([]Coord, 0)
					checking_num = true
				}
				num = num * 10 + int(s[x]) - int('0')
				lst = append(lst, Coord{x,y})
			} else {
				if checking_num {
					my_map[num] = append(my_map[num], lst)
					checking_num = false
				}
				if s[x] != '.' {
					symbols = append(symbols, Coord{x,y})
				}
			}
			x += 1

		}

		if checking_num {
			my_map[num] = append(my_map[num], lst)
			checking_num = false
		}

		y+=1
	}

	fmt.Println(symbols)

	for key, val2 := range my_map {
		for _, val := range val2 {
			close := false
			for _, loc := range val {
				x := loc.x
				y := loc.y
				for _, pos := range symbols {
					xs := pos.x
					ys := pos.y

					if (abs(x-xs) <= 1 && abs(y-ys) <= 1) {
						close = true
						break
					}
				}
				if close {
					break
				}
			}
			
			if close {
				fmt.Println(key)
				fmt.Println(val)
				sum += key
			}
		}
	}

	return sum
}

func part2(scanner *bufio.Scanner) int {
	sum := 0

	my_map := make(map[int]([][]Coord))
	symbols := make([]Coord, 0)

	y := 0
	for scanner.Scan() {
		s := scanner.Text()

		len := len(s)
		
		x := 0

		checking_num := true
		num := 0

		lst := make([]Coord, 0)

		for x < len {
			if s[x] >= '0' && s[x] <= '9' {
				if !checking_num {
					num = 0
					lst = make([]Coord, 0)
					checking_num = true
				}
				num = num * 10 + int(s[x]) - int('0')
				lst = append(lst, Coord{x,y})
			} else {
				if checking_num {
					my_map[num] = append(my_map[num], lst)
					checking_num = false
				}
				if s[x] == '*' {
					symbols = append(symbols, Coord{x,y})
				}
			}
			x += 1

		}

		if checking_num {
			my_map[num] = append(my_map[num], lst)
			checking_num = false
		}

		y+=1
	}

	for _, pos := range symbols {
		valid1, valid2 := -1, -1
		gear := true

		xs := pos.x
		ys := pos.y

		set_vals := func(val int) {
			if (valid1 == -1) {
				valid1 = val
			} else if (valid2 == -1) {
				valid2 = val
			} else {
				gear = false
			}
		}

		for key, val2 := range my_map {
			for _, val := range val2 {
				for _, loc := range val {
					x := loc.x
					y := loc.y

					if (abs(x-xs) <= 1 && abs(y-ys) <= 1) {
						set_vals(key)
						break
					}
				}
			}
		}

		fmt.Println(valid1)
		fmt.Println(valid2)

		if gear && valid1 != -1 && valid2 != -1 {
			sum += valid1 * valid2
		}
	}

	return sum
}

func main() {

	file, err := os.Open("input3.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    //sum := part1(scanner)
	// fmt.Println(sum)

	sum := part2(scanner)
	fmt.Println(sum)
}