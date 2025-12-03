package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func dial(number int, command string) (int, int) {
	command = strings.TrimSpace(command)
	if command == "" {
		return number, 0
	}
	if len(command) < 2 {
		log.Fatalf("invalid command %q", command)
	}

	dir := command[0]
	valueStr := command[1:]
	steps, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("invalid step value in %q: %v", command, err)
	}
	if steps < 0 {
		log.Fatalf("negative steps not supported: %q", command)
	}

	p := number
	var d int       // direction: +1 (R) or -1 (L)
	var tToZero int // steps until first time we hit 0 along this direction

	switch dir {
	case 'R':
		d = +1
		tToZero = (100 - p) % 100
		if tToZero == 0 {
			tToZero = 100
		}
	case 'L':
		d = -1
		tToZero = p
		if tToZero == 0 {
			tToZero = 100
		}
	default:
		log.Fatalf("unknown direction %q in command %q", string(dir), command)
	}

	// Count hits of 0 during this rotation.
	hits := 0
	if steps >= tToZero {
		hits = 1 + (steps-tToZero)/100
	}

	// New position after applying all steps.
	newPos := (p + d*steps) % 100
	if newPos < 0 {
		newPos += 100
	}

	return newPos, hits
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: %s <1|2> <file>", os.Args[0])
	}

	mode := os.Args[1]
	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading file: %v", err)
	}

	current := 50
	password := 0

	for _, line := range lines {
		next, hits := dial(current, line)

		switch mode {
		case "1":
			// Part 1: count only when the dial ends at 0 after a rotation.
			if next == 0 {
				password++
			}
		case "2":
			// Part 2: count *all clicks* that land on 0 (during + end).
			password += hits
		default:
			log.Fatalf("unknown mode %q (use 1 or 2)", mode)
		}

		current = next
	}

	fmt.Println("Password:", password)
}
