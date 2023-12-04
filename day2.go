package main

import (
    "fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func part1() int {
	sum := 0

	file, err := os.Open("input2.txt")

	red, green, blue := 12, 13, 14

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	gameNum := 1

	for scanner.Scan() {
		s := scanner.Text()
		
		ind := strings.Index(s, ":")
		s = s[ind + 1 :]

		list_str := strings.Split(s, ";")

		valid := true
		
		for _,substr := range list_str {
			new_list := strings.Split(substr, ",")

			for _,g := range new_list {
				if strings.Contains(g, "red") {
					str_red := g[1: strings.Index(g, "red") - 1]
					num_red, err := strconv.Atoi(str_red)
					if err != nil {
						// ... handle error
						panic(err)
					}
					if (num_red > red) {
						valid = false
					}
				}
				if strings.Contains(g, "green") {
					str_green := g[1: strings.Index(g, "green") - 1]
					num_green, err := strconv.Atoi(str_green)
					if err != nil {
						// ... handle error
						panic(err)
					}
					if (num_green > green) {
						valid = false
					}
				}
				if strings.Contains(g, "blue") {
					str_blue := g[1: strings.Index(g, "blue") - 1]
					num_blue, err := strconv.Atoi(str_blue)
					if err != nil {
						// ... handle error
						panic(err)
					}
					if (num_blue > blue) {
						valid = false
					}
				}
			}
		}
		
		if valid {
			sum += gameNum
		}

		gameNum += 1
	}

	return sum
}

func part2() int {
	sum := 0

	file, err := os.Open("input2.txt")


	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	gameNum := 1

	for scanner.Scan() {
		// s := scanner.Text()

		red, green, blue := 0,0,0
		
		s := scanner.Text()
		
		ind := strings.Index(s, ":")
		s = s[ind + 1 :]

		list_str := strings.Split(s, ";")
		
		for _, substr := range list_str {
			new_list := strings.Split(substr, ",")

			for _,g := range new_list {
				if strings.Contains(g, "red") {
					str_red := g[1: strings.Index(g, "red") - 1]
					num_red, err := strconv.Atoi(str_red)
					if err != nil {
						// ... handle error
						panic(err)
					}
					if (num_red > red) {
						red = num_red
					}
				}
				if strings.Contains(g, "green") {
					str_green := g[1: strings.Index(g, "green") - 1]
					num_green, err := strconv.Atoi(str_green)
					if err != nil {
						// ... handle error
						panic(err)
					}
					if (num_green > green) {
						green = num_green
					}
				}
				if strings.Contains(g, "blue") {
					str_blue := g[1: strings.Index(g, "blue") - 1]
					num_blue, err := strconv.Atoi(str_blue)
					if err != nil {
						// ... handle error
						panic(err)
					}
					if (num_blue > blue) {
						blue = num_blue
					}
				}
			}
		}
		
		sum += red * blue * green

		gameNum += 1
	}

	return sum
}

func main() {
    sum := part1()
	fmt.Println(sum)

	sum = part2()
	fmt.Println(sum)
}