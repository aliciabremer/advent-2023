package main

import (
    "fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func makeArray(scanner *bufio.Scanner) [][3]int {
	
	arr := make([][3]int, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if (s == "") {
			break
		}

		split_str := strings.Split(s, " ")
		var nums [3]int
		for i, num_str := range split_str {
			int_val, err := strconv.Atoi(num_str)
			if (err != nil) {
				panic(err)
			}
			nums[i] = int_val
		}

		arr = append(arr, nums)

	}
	return arr
}

func checkArray(num int, arr [][3]int) int {
	for _, ran := range arr {
		if (num - ran[1] >= 0 && num-ran[1] < ran[2]) {
			return ran[0] + num-ran[1]
		}
	}
	return num
}

func part1(scanner *bufio.Scanner) int {

	scanner.Scan() 
	s := scanner.Text()
	s = s[6:]
	
	seeds_str := strings.Split(s, " ")
	seeds := make([]int, 0)

	for _, num := range seeds_str {
		if num == "" || num == " " {
			continue
		}
		int_num, err := strconv.Atoi(num)
		if (err != nil) {
			panic(err)
		}
		seeds = append(seeds, int_num)
	}

	fmt.Println(seeds)
		
	

	maps := make([][][3]int, 0)

	scanner.Scan()

	for scanner.Scan() {
		maps = append(maps, makeArray(scanner))
	}

	fmt.Println(maps)

	
	min := -1

	for _, seed := range seeds {
		num := seed
		for _, m := range maps {
			num = checkArray(num, m)
		}

		if (min == -1 || num < min) {
			min = num
		}
	}
	
	return min
}

func max(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i int, j int) int {
	if i <= j {
		return i
	}
	return j
}

func ranges_contained(r1 [2]int, r2 [2]int) bool {
	if (r1[0] >= r2[0] && r1[0] < r2[0] + r2[1]) {
		return true
	}
	if (r2[0] >= r1[0] && r2[0] < r1[0] + r1[1]) {
		return true
	}

	return false
}

func checkArray2(r [2]int, arr [][3]int) [][2]int {
	// change this to split r into separate ranges of things
	// just do it recursively
	ranges_new := make([][2]int, 0)
	start_range := r[0]
	len_range := r[1]
	for _, ran := range arr {
		if (ranges_contained(r, [2]int{ran[1], ran[2]})) {
			diff_start_ran_start := max(start_range - ran[1], 0)
			new_start := ran[0] + diff_start_ran_start
			new_len := 0
			if ran[1] + ran[2] >= start_range + len_range {
				new_len = start_range + len_range - (ran[1] + diff_start_ran_start)
			} else {
				new_len = ran[1] + ran[2] - (ran[1] + diff_start_ran_start)
			}
			ranges_new = append(ranges_new, [2]int{new_start, new_len})
			if (start_range - ran[1] < 0) {
				ranges_new = append(ranges_new, checkArray2([2]int{start_range,  ran[1]- start_range}, arr)...)
			}
			if (start_range + len_range > ran[1] + ran[2]) {
				diff := start_range + len_range  - (ran[1] + ran[2])
				ranges_new = append(ranges_new, checkArray2([2]int{ran[1]+ran[2], diff}, arr)...)
			}

			return ranges_new
		}
	}
	return append(ranges_new, r)
}

func get_range_comb(r1 [2]int, r2 [2]int) [2]int {
	if (r1[0] >= r2[0]) {
		other_end := max(r1[0] + r1[1], r2[0] + r2[1])
		return [2]int {r2[0], other_end - r2[0] + 1}
	}
	other_end := max(r1[0] + r1[1], r2[0] + r2[1])
	return [2]int {r1[0], other_end - r1[0] + 1}
}

func part2(scanner *bufio.Scanner) int {
	scanner.Scan() 
	s := scanner.Text()
	s = s[6:]
	
	seeds_str := strings.Split(s, " ")
	seeds := make([]int, 0)

	for _, num := range seeds_str {
		if num == "" || num == " " {
			continue
		}
		int_num, err := strconv.Atoi(num)
		if (err != nil) {
			panic(err)
		}
		seeds = append(seeds, int_num)
	}
		
	

	maps := make([][][3]int, 0)

	scanner.Scan()

	for scanner.Scan() {
		maps = append(maps, makeArray(scanner))
	}

	fmt.Println(maps)

	

	ranges := make([][2]int, 0)

	for i := 0; i < len(seeds); i+=2 {
		var nums [2]int
		nums[0] = seeds[i]
		nums[1] = seeds[i+1]
		ranges = append(ranges, nums)
	}

	fmt.Println(ranges)

	for _, m := range maps {
		new_ranges := make([][2]int, 0)
		// fmt.Println("-------")
		for _, r := range ranges {
			// fmt.Println(r)
			new_arr := checkArray2(r, m)
			// fmt.Println(new_arr)
			for _, new_r := range new_arr {
				added := false
				for j, other_r := range new_ranges {
					if (ranges_contained(new_r, other_r)) {
						new_ranges[j] = get_range_comb(new_r, other_r)
						added = true
						break
					}
				}
				if !added {
					new_ranges = append(new_ranges, new_r)
				}
			}
			// fmt.Println(new_ranges)
		}
		ranges = new_ranges
		// fmt.Println(ranges)
	}

	min := -1

	for _, r := range ranges {
		if min == -1 {
			min = r[0]
		} else {
			if r[0] < min {
				min = r[0]
			}
		}
	}

	
	return min
}

func main() {

	filename := "input5.txt"

	// file, err := os.Open(filename)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer file.Close()

	// scanner := bufio.NewScanner(file)

    // sum := part1(scanner)
	// fmt.Println(sum)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := part2(scanner)
	fmt.Println(sum)
}