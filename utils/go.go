package aoc_utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func ElapsedRun(f func(string) int, input string) (int, time.Duration) {
	start := time.Now()
	res := f(input)
	elapsed := time.Since(start)
    
	return res, elapsed
}

func RunPart(part_id int, f func(string)int, in_path string)string{
    file, err := os.ReadFile(in_path)
    if err != nil {
        panic(err)
    }

    input := strings.TrimSpace(string(file))
    res, elapsed := ElapsedRun(f, input)
    return fmt.Sprintf("Part %d:\t{%d}\t Elapsed: %s\n",part_id, res, elapsed.String())
}
