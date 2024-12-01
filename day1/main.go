package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/radar07/aoc2024-go/utils"
)

func diff(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}

func Solve1() {
	result := utils.ReadFile("day1/input.txt")
	ans := 0

	left, right := make([]int, 0), make([]int, 0)

	for i := range len(result) {
		splits := strings.Split(result[i], " ")
		a, _ := strconv.Atoi(splits[0])
		b, _ := strconv.Atoi(splits[3])
		left = append(left, a)
		right = append(right, b)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := range len(result) {
		ans += diff(left[i], right[i])
	}

	fmt.Println("PART 1:", ans)
}

func Solve2() {
	result := utils.ReadFile("day1/input.txt")
	ans := 0

	list := make([]int, 0)
	freq := make(map[int]int)

	for i := range len(result) {
		splits := strings.Split(result[i], " ")
		a, _ := strconv.Atoi(splits[0])
		b, _ := strconv.Atoi(splits[3])

		list = append(list, a)
		freq[b]++
	}

	for i := range len(result) {
		ans += list[i] * freq[list[i]]
	}

	fmt.Println("PART 2:", ans)
}
