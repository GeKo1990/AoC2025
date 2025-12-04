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
	scanner := bufio.NewScanner(os.Stdin)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue // skip empty lines
		}
		row := []rune(line)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	fmt.Printf("result: %d", part1(grid))
}
