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

	// Part 1: auf einer Kopie arbeiten, damit das Original-Grid für Part 2 unverändert bleibt
	gridPart1 := cloneGrid(grid)
	splits := processGrid(gridPart1)
	fmt.Printf("\nSplits at ^: %d\n", splits)

	// Part 2: Timelines im Original-Grid zählen
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

// cloneGrid erstellt eine tiefe Kopie des Grids.
func cloneGrid(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))
	for i := range src {
		dst[i] = make([]byte, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

// Part 1: klassische Strahl-Simulation, zählt wie oft ein Beam an einem '^' gesplittet wird.
func processGrid(grid [][]byte) int {
	splitCount := 0

	for row := 1; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			elementAbove := grid[row-1][col]
			currentElement := grid[row][col]

			// Beam fällt von S oder | nach unten in einen Punkt .
			if (elementAbove == 'S' || elementAbove == '|') && currentElement == '.' {
				grid[row][col] = '|'
				currentElement = '|'
			}

			// Splitter: nur wenn von oben ein Beam kommt
			if grid[row][col] == '^' && (elementAbove == 'S' || elementAbove == '|') {
				splitCount++

				// neue Beams nach links/rechts, wenn dort noch leer ist
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

// Part 2: Quantum-Version – wie viele verschiedene Timelines gibt es?
// Hier arbeiten wir nicht mit | im Grid, sondern nur mit Zählern pro Spalte.
func countTimelines(grid [][]byte) int64 {
	h := len(grid)
	if h == 0 {
		return 0
	}
	w := len(grid[0])

	// Startposition S finden
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

	// paths[c] = wie viele Timelines laufen aktuell in Spalte c nach unten
	paths := make([]int64, w)
	paths[sc] = 1 // direkt unter S startet genau 1 Timeline

	// wir gehen Zeile für Zeile unterhalb von S nach unten
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
				// Splitter: jede Timeline teilt sich in eine links + eine rechts
				if c > 0 {
					next[c-1] += k
				}
				if c+1 < w {
					next[c+1] += k
				}
				// nach unten geht nichts durch den Splitter

			default:
				// alles andere ('.', 'S', usw.) → Timeline läuft nach unten weiter
				next[c] += k
			}
		}

		paths = next
	}

	// am unteren Rand: alle Timelines fallen aus dem Manifold heraus
	var total int64
	for _, k := range paths {
		total += k
	}
	return total
}
