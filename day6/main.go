package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Column struct {
	Values []int
	Op     byte
}

func evalColumn(c Column) int {
	if len(c.Values) == 0 {
		return 0
	}

	switch c.Op {
	case '*':
		res := 1
		for _, v := range c.Values {
			res *= v
		}
		return res
	case '+':
		res := 0
		for _, v := range c.Values {
			res += v
		}
		return res
	case '-':
		res := 0
		for _, v := range c.Values {
			res -= v
		}
		return res
	default:
		return 0
	}
}

func part1(cols []Column) int {
	sum := 0
	for _, c := range cols {
		sum += evalColumn(c)
	}
	return sum
}

func buildColumnsForPart1(numLines []string, opLine string) []Column {
	firstRow := strings.Fields(numLines[0])
	cols := make([]Column, len(firstRow))

	for _, row := range numLines {
		parts := strings.Fields(row)
		for i, p := range parts {
			v, _ := strconv.Atoi(p)
			cols[i].Values = append(cols[i].Values, v)
		}
	}

	opTokens := strings.Fields(opLine)
	for i, op := range opTokens {
		if i < len(cols) && len(op) > 0 {
			cols[i].Op = op[0]
		}
	}
	return cols
}

func part2(numLines []string, opLine string) int {
	maxLen := len(opLine)
	for _, l := range numLines {
		if len(l) > maxLen {
			maxLen = len(l)
		}
	}

	lines := make([]string, len(numLines))
	for i, l := range numLines {
		if len(l) < maxLen {
			l += strings.Repeat(" ", maxLen-len(l))
		}
		lines[i] = l
	}
	if len(opLine) < maxLen {
		opLine += strings.Repeat(" ", maxLen-len(opLine))
	}

	var ops []byte
	for i := maxLen - 1; i >= 0; i-- {
		ch := opLine[i]
		if ch == '+' || ch == '*' || ch == '-' {
			ops = append(ops, ch)
		}
	}

	allSpacesAt := func(col int) bool {
		for _, row := range lines {
			if row[col] != ' ' {
				return false
			}
		}
		return true
	}

	var problems []Column
	var numbers []int
	opIndex := 0

	for col := maxLen - 1; col >= -1; col-- {
		if col == -1 || allSpacesAt(col) {
			if len(numbers) > 0 && opIndex < len(ops) {
				vals := make([]int, len(numbers))
				copy(vals, numbers)

				problems = append(problems, Column{
					Values: vals,
					Op:     ops[opIndex],
				})
				opIndex++
				numbers = numbers[:0]
			}
		} else {
			numStr := ""
			for _, row := range lines {
				ch := row[col]
				if ch != ' ' {
					numStr += string(ch)
				}
			}
			numStr = strings.TrimSpace(numStr)
			if numStr != "" {
				v, _ := strconv.Atoi(numStr)
				numbers = append(numbers, v)
			}
		}
	}

	return part1(problems)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [1|2]")
		return
	}
	mode := os.Args[1]

	sc := bufio.NewScanner(os.Stdin)
	var raw []string
	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) != "" {
			raw = append(raw, line)
		}
	}
	if len(raw) < 2 {
		fmt.Println("need at least one number row and one operator row")
		return
	}

	opLine := raw[len(raw)-1]
	numLines := raw[:len(raw)-1]

	switch mode {
	case "1":
		cols := buildColumnsForPart1(numLines, opLine)
		fmt.Println(part1(cols))
	case "2":
		fmt.Println(part2(numLines, opLine))
	default:
		fmt.Println("Usage: go run main.go [1|2]")
	}
}
