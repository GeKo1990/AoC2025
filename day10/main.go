package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	target  []bool  // desired indicator state (false='.', true='#')
	buttons [][]int // each button toggles these indices
	jolts   []int   // currently unused
}

func parseMachine(line string) *Machine {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}

	m := &Machine{}

	lb := strings.Index(line, "[")
	rb := strings.Index(line, "]")
	if lb == -1 || rb == -1 || rb <= lb {
		return nil
	}
	pattern := line[lb+1 : rb]
	for _, ch := range pattern {
		switch ch {
		case '.':
			m.target = append(m.target, false)
		case '#':
			m.target = append(m.target, true)
		}
	}

	var joltsStr string
	lc := strings.Index(line, "{")
	rc := strings.Index(line, "}")
	if lc != -1 && rc != -1 && rc > lc {
		joltsStr = line[lc+1 : rc]
		for _, part := range strings.Split(joltsStr, ",") {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			v, err := strconv.Atoi(part)
			if err == nil {
				m.jolts = append(m.jolts, v)
			}
		}
	}

	buttonsPart := line
	if lc != -1 {
		buttonsPart = line[rb+1 : lc]
	}
	s := buttonsPart
	for {
		lp := strings.Index(s, "(")
		if lp == -1 {
			break
		}
		rp := strings.Index(s[lp+1:], ")")
		if rp == -1 {
			break
		}
		rp += lp + 1
		inner := strings.TrimSpace(s[lp+1 : rp])
		if inner != "" {
			var btn []int
			for _, p := range strings.Split(inner, ",") {
				p = strings.TrimSpace(p)
				if p == "" {
					continue
				}
				v, err := strconv.Atoi(p)
				if err == nil {
					btn = append(btn, v)
				}
			}
			if len(btn) > 0 {
				m.buttons = append(m.buttons, btn)
			}
		}
		s = s[rp+1:]
	}

	return m
}

func minPresses(m *Machine) int {
	// encode target as bitmask
	var target uint64
	for i, on := range m.target {
		if on {
			target |= 1 << uint(i)
		}
	}

	B := len(m.buttons)
	if B == 0 {
		// no buttons: either already satisfied or impossible
		if target == 0 {
			return 0
		}
		return -1
	}

	// encode each button as bitmask
	btnMasks := make([]uint64, B)
	for i, btn := range m.buttons {
		var mask uint64
		for _, idx := range btn {
			mask |= 1 << uint(idx)
		}
		btnMasks[i] = mask
	}

	best := -1
	// enumerate all subsets of buttons
	total := 1 << uint(B)
	for subset := 0; subset < total; subset++ {
		var state uint64
		presses := 0
		for b := 0; b < B; b++ {
			if subset&(1<<uint(b)) != 0 {
				state ^= btnMasks[b]
				presses++
			}
		}
		if state == target {
			if best == -1 || presses < best {
				best = presses
			}
		}
	}
	return best
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	var machines []*Machine

	for in.Scan() {
		line := in.Text()
		m := parseMachine(line)
		if m != nil {
			machines = append(machines, m)
		}
	}
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
	}

	for i, m := range machines {
		fmt.Printf("Machine %d:\n", i+1)
		fmt.Printf("  target  (%d lights): ", len(m.target))
		for _, b := range m.target {
			if b {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()

		fmt.Printf("  buttons (%d):\n", len(m.buttons))
		for bi, btn := range m.buttons {
			fmt.Printf("    %d: %v\n", bi, btn)
		}

		fmt.Printf("  jolts   (%d): %v\n", len(m.jolts), m.jolts)
	}

	totalPresses := 0
	for _, m := range machines {
		best := minPresses(m)
		if best < 0 {
			fmt.Fprintln(os.Stderr, "machine not solvable?")
			continue
		}
		totalPresses += best
	}
	fmt.Println("total Presses:", totalPresses)
}
