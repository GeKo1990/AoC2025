package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type P struct{ x, y int }

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// checks if the rectangle [xmin,xmax] x [ymin,ymax] is fully inside
// the orthogonal polygon defined by pts
func rectInside(xmin, xmax, ymin, ymax int, pts []P) bool {
	// point-in-polygon check for the rectangle center
	cx := float64(xmin+xmax) / 2.0
	cy := float64(ymin+ymax) / 2.0

	inside := false
	n := len(pts)
	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		xi, yi := float64(pts[i].x), float64(pts[i].y)
		xj, yj := float64(pts[j].x), float64(pts[j].y)

		if (yi > cy) != (yj > cy) {
			xint := xi + (xj-xi)*(cy-yi)/(yj-yi)
			if xint > cx {
				inside = !inside
			}
		}
	}
	if !inside {
		return false
	}

	// ensure no polygon edge crosses the rectangle's *interior*
	for i := 0; i < n; i++ {
		x1, y1 := pts[i].x, pts[i].y
		x2, y2 := pts[(i+1)%n].x, pts[(i+1)%n].y

		if x1 == x2 { // vertical edge
			xv := x1
			if xv <= xmin || xv >= xmax {
				continue
			}
			ya, yb := y1, y2
			if ya > yb {
				ya, yb = yb, ya
			}
			lo := max(ya, ymin)
			hi := min(yb, ymax)
			if hi > lo { // overlaps interior in y
				return false
			}
		} else if y1 == y2 { // horizontal edge
			yv := y1
			if yv <= ymin || yv >= ymax {
				continue
			}
			xa, xb := x1, x2
			if xa > xb {
				xa, xb = xb, xa
			}
			lo := max(xa, xmin)
			hi := min(xb, xmax)
			if hi > lo { // overlaps interior in x
				return false
			}
		}
	}

	return true
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	var pts []P
	for in.Scan() {
		line := strings.TrimSpace(in.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			continue
		}
		pts = append(pts, P{x, y})
	}
	if len(pts) < 2 {
		fmt.Println(0)
		fmt.Println(0)
		return
	}

	n := len(pts)

	// Part 1: no constraints, just red corners
	part1 := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := abs(pts[i].x - pts[j].x)
			dy := abs(pts[i].y - pts[j].y)
			area := (dx + 1) * (dy + 1)
			if area > part1 {
				part1 = area
			}
		}
	}

	// Part 2: rectangle must lie fully within red+green region (polygon)
	part2 := 0
	for i := 0; i < n; i++ {
		x1, y1 := pts[i].x, pts[i].y
		for j := i + 1; j < n; j++ {
			x2, y2 := pts[j].x, pts[j].y

			xmin, xmax := x1, x2
			if xmin > xmax {
				xmin, xmax = xmax, xmin
			}
			ymin, ymax := y1, y2
			if ymin > ymax {
				ymin, ymax = ymax, ymin
			}

			area := (xmax - xmin + 1) * (ymax - ymin + 1)
			if area <= part2 {
				continue
			}

			if rectInside(xmin, xmax, ymin, ymax, pts) {
				part2 = area
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
