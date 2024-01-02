package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "johnbakhmat.tech/m/v2/aoc/utils"
)

func main() {
	fmt.Println(utils.RunPart(1, Part1, "./input.1.txt"))
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	gamma := make([]rune, width)
	epsilon := make([]rune, width)

	for i := 0; i < width; i++ {
		zero_count := 0
		for _, line := range lines {
			if line[i] == '0' {
				zero_count++
			}
		}
        if zero_count > len(lines)/2{
            gamma[i] = '0'
            epsilon[i] = '1'
        } else {
            gamma[i] = '1'
            epsilon[i] = '0'
        }
	}

	gamma_parsed, err := strconv.ParseInt(string(gamma), 2, 64)
	if err != nil {
		panic(err)
	}

	epsilon_parsed, err := strconv.ParseInt(string(epsilon), 2, 64)
	if err != nil {
		panic(err)
	}

    return int(epsilon_parsed * gamma_parsed)


}
