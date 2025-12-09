package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := readGrid()

	fmt.Println("Grid loaded:")
	for _, row := range grid {
		fmt.Println(string(row))
	}

	gridPart1 := cloneGrid(grid)
	splits := processGrid(gridPart1)
	fmt.Printf("\nSplits at ^: %d\n", splits)

	timelines := countTimelines(grid)
	fmt.Printf("Timelines: %d\n", timelines)
}

func readGrid() [][]byte {
	sc := bufio.NewScanner(os.Stdin)
	var grid [][]byte

	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []byte(line))
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	return grid
}

func cloneGrid(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))
	for i := range src {
		dst[i] = make([]byte, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func processGrid(grid [][]byte) int {
	splitCount := 0

	for row := 1; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			elementAbove := grid[row-1][col]
			currentElement := grid[row][col]

			if (elementAbove == 'S' || elementAbove == '|') && currentElement == '.' {
				grid[row][col] = '|'
				currentElement = '|'
			}

			if grid[row][col] == '^' && (elementAbove == 'S' || elementAbove == '|') {
				splitCount++

				if col > 0 && grid[row][col-1] == '.' {
					grid[row][col-1] = '|'
				}
				if col+1 < len(grid[row]) && grid[row][col+1] == '.' {
					grid[row][col+1] = '|'
				}
			}
		}
	}

	fmt.Println("\nGrid after processing (Part 1):")
	for _, row := range grid {
		fmt.Println(string(row))
	}

	return splitCount
}

func countTimelines(grid [][]byte) int64 {
	h := len(grid)
	if h == 0 {
		return 0
	}
	w := len(grid[0])

	sr, sc := -1, -1
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] == 'S' {
				sr, sc = r, c
				break
			}
		}
		if sr != -1 {
			break
		}
	}
	if sr == -1 {
		return 0
	}

	paths := make([]int64, w)
	paths[sc] = 1

	for r := sr + 1; r < h; r++ {
		next := make([]int64, w)

		for c := 0; c < w; c++ {
			k := paths[c]
			if k == 0 {
				continue
			}

			ch := grid[r][c]

			switch ch {
			case '^':
				if c > 0 {
					next[c-1] += k
				}
				if c+1 < w {
					next[c+1] += k
				}

			default:
				next[c] += k
			}
		}

		paths = next
	}

	var total int64
	for _, k := range paths {
		total += k
	}
	return total
}
