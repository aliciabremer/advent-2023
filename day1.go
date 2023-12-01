package main

import (
    "fmt"
	"bufio"
	"os"
)

func part1() int {
	sum := 0

	file, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		var first, last int
		for i := 0; i < len(s); i++ {
			char := int(s[i])
			if char >= int('0') && char <= ('9') {
				if first == 0 {
					first = char - int('0')
				}
				last = char - int('0')
			}
		}
		concat := first* 10 + last
		sum += concat
	}

	return sum
}

func part2() int {
	sum := 0

	file, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		var first, last int

		set_vals := func(val int) {
			if (first == 0) {
				first = val
			}
			last = val
		}

		for i := 0; i < len(s); i++ {
			char := int(s[i])
			if char >= int('0') && char <= ('9') {
				set_vals(char - int('0'))
			} else {
				if (i + 2 < len(s)) {
					if (s[i:i+3] == "one") {
						set_vals(1)
					}
					if (s[i:i+3]=="two") {
						set_vals(2)
					}
					if (s[i:i+3] == "six") {
						set_vals(6)
					}
				}
				if (i + 3 < len(s)) {
					if (s[i:i+4] == "four") {
						set_vals(4)
					}
					if (s[i:i+4] == "five") {
						set_vals(5)
					}
					if (s[i:i+4] == "nine") {
						set_vals(9)
					}
				}
				if (i + 4 < len(s)) {
					if (s[i:i+5] == "three") {
						set_vals(3)
					}
					if (s[i:i+5] == "seven") {
						set_vals(7)
					}
					if (s[i:i+5] == "eight") {
						set_vals(8)
					}
				}
			}
			
		}
		concat := first* 10 + last
		sum += concat
	}

	return sum
}

func main() {
    sum := part1()
	fmt.Println(sum)

	sum = part2()
	fmt.Println(sum)
}