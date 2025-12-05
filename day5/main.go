package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start, end int64
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <part 1|2> < input.txt")
		return
	}

	part := os.Args[1]

	ranges, numbers, err := readInput(os.Stdin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch part {
	case "1":
		fmt.Println("Part 1:", solvePart1(ranges, numbers))
	case "2":
		fmt.Println("Part 2:", solvePart2(ranges))
	default:
		fmt.Println("Unknown part:", part)
	}
}

func readInput(r io.Reader) ([]interval, []int64, error) {
	sc := bufio.NewScanner(r)

	var ranges []interval
	var numbers []int64
	mode := "ranges"

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())

		if line == "" {
			mode = "numbers"
			continue
		}

		if mode == "ranges" {
			parts := strings.SplitN(line, "-", 2)
			start, _ := strconv.ParseInt(parts[0], 10, 64)
			end, _ := strconv.ParseInt(parts[1], 10, 64)
			if end < start {
				start, end = end, start
			}
			ranges = append(ranges, interval{start, end})

		} else {
			n, _ := strconv.ParseInt(line, 10, 64)
			numbers = append(numbers, n)
		}
	}

	return ranges, numbers, sc.Err()
}

func solvePart1(ranges []interval, numbers []int64) int64 {
	var count int64

	for _, n := range numbers {
		for _, r := range ranges {
			if n >= r.start && n <= r.end {
				count++
				break
			}
		}
	}
	return count
}

func solvePart2(ranges []interval) int64 {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	// Merge overlapping intervals
	curStart := ranges[0].start
	curEnd := ranges[0].end
	var total int64

	for _, r := range ranges[1:] {
		if r.start <= curEnd+1 {
			if r.end > curEnd {
				curEnd = r.end
			}
		} else {
			total += curEnd - curStart + 1
			curStart = r.start
			curEnd = r.end
		}
	}

	total += curEnd - curStart + 1

	return total
}
