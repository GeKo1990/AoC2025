package main

import (
	"bufio"
	"fmt"
	"os"
)

func bestK(line string, k int) int64 {
	digits := make([]int, 0, len(line))
	for _, ch := range line {
		if ch >= '0' && ch <= '9' {
			digits = append(digits, int(ch-'0'))
		}
	}

	n := len(digits)
	if n < k || k <= 0 {
		return 0
	}

	toDrop := n - k
	stack := make([]int, 0, n)

	for _, d := range digits {
		for toDrop > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1]
			toDrop--
		}
		stack = append(stack, d)
	}

	if len(stack) > k {
		stack = stack[:k]
	}

	var result int64
	for _, d := range stack {
		result = result*10 + int64(d)
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go run main.go <1|2> < input.txt>")
		return
	}

	mode := os.Args[1]
	k := 0
	if mode == "1" {
		k = 2
	} else if mode == "2" {
		k = 12
	} else {
		fmt.Fprintln(os.Stderr, "unknown mode:", mode, "(use 1 or 2)")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	var total int64

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		total += bestK(line, k)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		return
	}

	fmt.Println(total)
}
