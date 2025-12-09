package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// We conceptually follow Kruskal's algorithm with a Disjoint Set (Union-Find):
// - Kruskal:      https://en.wikipedia.org/wiki/Kruskal%27s_algorithm
// - Disjoint set: https://en.wikipedia.org/wiki/Disjoint-set_data_structure

type Point3D struct {
	X, Y, Z int
}

type Edge struct {
	A, B  int   // indices of points
	Dist2 int64 // squared distance between them
}

type UnionFind struct {
	parent []int
	size   []int
	comps  int // number of connected components
}

func newUnionFind(n int) *UnionFind {
	p := make([]int, n)
	s := make([]int, n)
	for i := range p {
		p[i] = i
		s[i] = 1
	}
	return &UnionFind{parent: p, size: s, comps: n}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(a, b int) bool {
	ra := uf.find(a)
	rb := uf.find(b)
	if ra == rb {
		return false
	}
	// union by size
	if uf.size[ra] < uf.size[rb] {
		ra, rb = rb, ra
	}
	uf.parent[rb] = ra
	uf.size[ra] += uf.size[rb]
	uf.comps--
	return true
}

func (uf *UnionFind) componentSizes() []int {
	m := make(map[int]int)
	for i := range uf.parent {
		root := uf.find(i)
		m[root]++
	}
	sizes := make([]int, 0, len(m))
	for _, v := range m {
		sizes = append(sizes, v)
	}
	return sizes
}

func main() {
	points := readPoints3D()
	n := len(points)
	if n == 0 {
		fmt.Println("no points")
		return
	}

	// 1) generate all pairwise edges with squared distances
	var edges []Edge
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := int64(points[i].X - points[j].X)
			dy := int64(points[i].Y - points[j].Y)
			dz := int64(points[i].Z - points[j].Z)
			dist2 := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{A: i, B: j, Dist2: dist2})
		}
	}

	// 2) sort edges by distance ascending
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Dist2 < edges[j].Dist2
	})

	uf := newUnionFind(n)

	// ---------- Part 1 ----------
	const connectionsToConsider = 1000
	limit := connectionsToConsider
	if limit > len(edges) {
		limit = len(edges)
	}

	for i := 0; i < limit; i++ {
		e := edges[i]
		// even if this doesn't merge new components, it's still one of the "K shortest"
		uf.union(e.A, e.B)
	}

	sizes := uf.componentSizes()
	if len(sizes) < 3 {
		fmt.Println("Part 1: not enough components to take top 3, sizes:", sizes)
	} else {
		sort.Ints(sizes)
		last := len(sizes) - 1
		a, b, c := sizes[last], sizes[last-1], sizes[last-2]
		product := a * b * c
		fmt.Printf("Part 1: largest component sizes %d, %d, %d -> product = %d\n", a, b, c, product)
	}

	// ---------- Part 2 ----------
	// Continue from the same uf state.
	// We now keep going through the remaining edges and track the LAST edge
	// that actually merges two different components. When uf.comps == 1,
	// that last merging edge is the one we need.
	var lastMerge Edge
	for i := limit; i < len(edges) && uf.comps > 1; i++ {
		e := edges[i]
		if uf.union(e.A, e.B) {
			lastMerge = e
		}
	}

	if uf.comps != 1 {
		fmt.Println("Part 2: still more than one component, something is off")
		return
	}

	pa := points[lastMerge.A]
	pb := points[lastMerge.B]
	part2 := int64(pa.X) * int64(pb.X)

	fmt.Printf("Part 2: last connection between (%d,%d,%d) and (%d,%d,%d)\n",
		pa.X, pa.Y, pa.Z, pb.X, pb.Y, pb.Z)
	fmt.Printf("Part 2: X product = %d\n", part2)
}

func readPoints3D() []Point3D {
	scanner := bufio.NewScanner(os.Stdin)
	var points []Point3D

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			fmt.Fprintf(os.Stderr, "invalid line: %s\n", line)
			continue
		}

		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))

		points = append(points, Point3D{X: x, Y: y, Z: z})
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	return points
}
