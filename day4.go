package main

import (
    "fmt"
	"bufio"
	"os"
	"strings"
	// "strconv"
)


func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func part1(scanner *bufio.Scanner) int {
	sum := 0

	for scanner.Scan() {
		s := scanner.Text()

		val := 0

		ind := strings.Index(s, ":")
		s = s[ind + 1 :]

		ind = strings.Index(s, "|")
		win := s[0: ind]
		act := s[ind+1:]

		winning_num := strings.Split(win, " ")
		act_num := strings.Split(act, " ")


		for _, n := range act_num {
			num_n := strings.TrimSpace(n)


			if num_n == "" {
				continue
			}

			contains := false
			for _, w := range winning_num {
				num_w  := strings.TrimSpace(w)
				if num_w == num_n {
					contains = true
					break
				}
			}

			if contains {
				if val == 0 {
					val = 1
				} else {
					val *= 2
				}
			}
		}

		sum += val
	}


	return sum
}

func part2(scanner *bufio.Scanner) int {
	sum := 0

	num_matching := make([]int, 0)
	count_cards := make([]int, 0)

	for scanner.Scan() {
		s := scanner.Text()


		ind := strings.Index(s, ":")
		s = s[ind + 1 :]

		ind = strings.Index(s, "|")
		win := s[0: ind]
		act := s[ind+1:]

		winning_num := strings.Split(win, " ")
		act_num := strings.Split(act, " ")


		count := 0
		for _, n := range act_num {
			num_n := strings.TrimSpace(n)


			if num_n == "" {
				continue
			}

			contains := false
			for _, w := range winning_num {
				num_w  := strings.TrimSpace(w)
				if num_w == num_n {
					contains = true
					break
				}
			}

			if contains {
				count += 1
			}
		}

		num_matching = append(num_matching, count)
		count_cards = append(count_cards, 0)
	}


	for i, n := range count_cards {
		match := num_matching[i]


		for j := 0; j < match; j++ {
			count_cards[i+j+1] += n+1
		}

		sum += n + 1
	}

	return sum
}

func main() {

	filename := "input4.txt"

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    sum := part1(scanner)
	fmt.Println(sum)

	file, err = os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)

	sum = part2(scanner)
	fmt.Println(sum)
}