package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Run(f func(string) int, input string) (int, time.Duration) {
	start := time.Now()
	res := f(input)
	elapsed := time.Since(start)
	return res, elapsed
}

func main() {
	file, err := os.ReadFile("./input.1.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(file))

	res, elapsed := Run(Part1, input)
	fmt.Printf("Part 1:\t{%d}\t Elapsed: %dms\n", res, elapsed.Milliseconds())

	file, err = os.ReadFile("./input.2.txt")
	if err != nil {
		panic(err)
	}

	input = strings.TrimSpace(string(file))

	res, elapsed = Run(Part2, input)
	fmt.Printf("Part 2:\t{%d}\t Elapsed: %dms\n", res, elapsed.Milliseconds())

}
func Part1(input string) int {
	lines := strings.Split(input, "\n")

	count := 0
	prev := 90000
	for _, v := range lines {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("Couldn't parse int: %s\n %s", v, err))
		}
		if prev < num {
			count++
		}
		prev = num
	}
	return count
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	numbers := make([]int, len(lines))
	for i, v := range lines {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numbers[i] = num
	}

	group_sums := make([]int, len(lines)*3)

	for i := 0; i < len(lines)-2; i++ {
		end := clamp(i+3, len(lines), 0)
		group_sums[i] = sum(numbers[i:end])
	}

	count := 0
	prev := math.MaxInt
	for _, v := range group_sums {
		if prev < v {
			count++
		}
		prev = v
	}
	return count

}
func clamp(v int, max int, min int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func sum(arr []int) int {
	s := 0

	for _, v := range arr {
		s += v
	}
	return s
}
