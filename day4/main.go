package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(grid [][]rune) int {
	hits := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			neigh := countNeighbourScrolls(grid, y, x)
			if grid[y][x] == '@' {
				if neigh <= 3 {
					hits++
				}
			}
		}
	}

	return hits
}

func part2Step(grid [][]rune) (int, [][]rune) {
	marked := make([][]rune, len(grid))
	for i := range grid {
		marked[i] = make([]rune, len(grid[i]))
		copy(marked[i], grid[i])
	}

	hits := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '@' {
				continue
			}
			neigh := countNeighbourScrolls(grid, y, x)
			if neigh <= 3 {
				hits++
				marked[y][x] = 'x'
			}
		}
	}

	return hits, marked
}

func part2(grid [][]rune) (int, [][]rune) {
	hits, marked := part2Step(grid)

	if hits == 0 {
		return 0, grid
	}

	for y := range marked {
		for x := range marked[y] {
			if marked[y][x] == 'x' {
				marked[y][x] = '.'
			}
		}
	}

	restHits, finalGrid := part2(marked)

	return hits + restHits, finalGrid
}

func countNeighbourScrolls(grid [][]rune, y, x int) int {
	h := len(grid)
	w := len(grid[0])
	count := 0

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, d := range dirs {
		ny := y + d[0]
		nx := x + d[1]

		if ny < 0 || ny >= h || nx < 0 || nx >= w {
			continue
		}

		if grid[ny][nx] == '@' {
			count++
		}
	}

	return count
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go run main.go <1|2> < input.txt>")
		os.Exit(1)
	}

	mode := os.Args[1]

	scanner := bufio.NewScanner(os.Stdin)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		row := []rune(line)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	switch mode {
	case "1":
		res := part1(grid)
		fmt.Printf("result: %d\n", res)

	case "2":
		res, _ := part2(grid)
		fmt.Printf("result: %d\n", res)

	default:
		fmt.Fprintf(os.Stderr, "unknown mode: %s (use 1 or 2)\n", mode)
		os.Exit(1)
	}
}
