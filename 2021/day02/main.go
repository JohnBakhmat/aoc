package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "johnbakhmat.tech/m/v2/aoc/utils"
)

func main() {
	fmt.Println(utils.RunPart(1, Part1, "./input.1.txt"))
	fmt.Println(utils.RunPart(2, Part2, "./input.2.txt"))
}

type Coord struct {
	x   int
	y   int
	aim int
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	coords := Coord{
		x: 0,
		y: 0,
	}

	for _, v := range lines {
		split := strings.Split(v, " ")
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		switch split[0] {
		case "forward":
			coords.x += amount
			break
		case "up":
			coords.y -= amount
			break
		case "down":
			coords.y += amount
			break
		}
	}
	return coords.x * coords.y
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	coords := Coord{
		x:   0,
		y:   0,
		aim: 0,
	}

	for _, v := range lines {
		split := strings.Split(v, " ")
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		switch split[0] {
		case "up":
			coords.aim -= amount
            break
		case "down":
			coords.aim += amount
			break
		case "forward":
			coords.x += amount
			coords.y += coords.aim * amount
			break
		}
	}

	return coords.y * coords.x
}
