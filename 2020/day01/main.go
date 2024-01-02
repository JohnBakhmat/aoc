package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "johnbakhmat.tech/m/v2/aoc/utils"
)


func main(){
    fmt.Println(utils.RunPart(1,Part1,"./input.1.txt"))
    fmt.Println(utils.RunPart(2,Part2,"./input.2.txt"))
}


func Part1(input string) int {
    target := 2020
	lines := strings.Split(input, "\n")
    nums := make([]int, len(lines))
    for i,v := range lines {
        int, err := strconv.Atoi(v)
        if err != nil{
            panic("Cant parse int")
        }
        nums[i] = int
    }

    hashMap := make(map[int]int, len(nums))

    for i,v := range(nums){
        t := target - v
        value, exists := hashMap[t]
        if (exists){
            return nums[value] * nums[i]
        }
        hashMap[v] = i
    }
	return 0
}

func Part2(input string) int {
	 target := 2020
	lines := strings.Split(input, "\n")
    nums := make([]int, len(lines))
    for i,v := range lines {
        int, err := strconv.Atoi(v)
        if err != nil{
            panic("Cant parse int")
        }
        nums[i] = int
    }

    for i:=0; i<len(nums); i++{
        for j := i+1; j<len(nums);j++{
            for k:= j+1; k<len(nums);k++{
                sum := nums[i] + nums[j] + nums[k]
                if sum == target {
                    return nums[i] * nums[j] * nums[k]
                }
            }
        } 
    }
    return 0

}
