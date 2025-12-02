package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Part 1: exactly "some block repeated twice"
func isDoublePatternExactlyTwice(n int64) bool {
	s := strconv.FormatInt(n, 10)
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}

// Part 2: "some block repeated at least twice"
func isRepeatedPatternAtLeastTwice(n int64) bool {
	s := strconv.FormatInt(n, 10)
	L := len(s)

	for size := 1; size <= L/2; size++ {
		if L%size != 0 {
			continue
		}
		times := L / size
		if times < 2 {
			continue
		}

		block := s[:size]
		ok := true
		for i := size; i < L; i += size {
			if s[i:i+size] != block {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go run main.go <1|2> < input.txt>")
		return
	}

	mode := os.Args[1]

	var isInvalid func(int64) bool
	switch mode {
	case "1":
		isInvalid = isDoublePatternExactlyTwice
	case "2":
		isInvalid = isRepeatedPatternAtLeastTwice
	default:
		fmt.Fprintln(os.Stderr, "unknown mode:", mode, "(use 1 or 2)")
		return
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		return
	}

	input := strings.TrimSpace(string(data))
	if input == "" {
		fmt.Fprintln(os.Stderr, "no input")
		return
	}

	ranges := strings.FieldsFunc(input, func(r rune) bool {
		return r == ','
	})

	var total int64

	for _, r := range ranges {
		if r == "" {
			continue
		}

		parts := strings.SplitN(r, "-", 2)
		if len(parts) != 2 {
			fmt.Fprintln(os.Stderr, "invalid range:", r)
			return
		}

		start, err1 := strconv.ParseInt(parts[0], 10, 64)
		end, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 != nil || err2 != nil {
			fmt.Fprintln(os.Stderr, "parse error in range:", r)
			return
		}

		if start > end {
			start, end = end, start
		}

		for x := start; x <= end; x++ {
			if isInvalid(x) {
				total += x
			}
		}
	}

	fmt.Println(total)
}
