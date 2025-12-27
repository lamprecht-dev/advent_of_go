package main

import (
	"aog/internal/aogutils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result := solve(data, 1000, 1)
	result2 := solve(data, 10000000, 2)
	fmt.Println(result, result2)
}

type vert struct {
	X float64
	Y float64
	Z float64
	// Disjoint part
	Parent *vert
	Size   int
}

func (v *vert) GetRoot() *vert {
	if v.Parent != v {
		v.Parent = v.Parent.GetRoot() // With path compression
	}
	return v.Parent
}

type edge struct {
	Vert1 *vert
	Vert2 *vert
	Dist  float64
}

func toFloat(val string) float64 {
	i, _ := strconv.Atoi(val)
	return float64(i)
}

func dist(v1, v2 vert) float64 {
	dx, dy, dz := v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2) + math.Pow(dz, 2))
}

func PrintComponents(boxes []*vert) {
	// Map from root -> list of verts in that component
	components := make(map[*vert][]*vert)

	for _, v := range boxes {
		root := v.GetRoot()
		components[root] = append(components[root], v)
	}

	fmt.Println("Disjoint set components:")

	for root, verts := range components {
		fmt.Printf("Root at (%.2f, %.2f, %.2f), Size %d:\n", root.X, root.Y, root.Z, root.Size)
		for _, v := range verts {
			fmt.Printf("  Vert: (%.2f, %.2f, %.2f)\n", v.X, v.Y, v.Z)
		}
		fmt.Println()
	}
}

func PrintEdge(e edge) {
	fmt.Printf("Edge: (%.2f, %.2f, %.2f) ↔ (%.2f, %.2f, %.2f), Dist: %.4f\n",
		e.Vert1.X, e.Vert1.Y, e.Vert1.Z,
		e.Vert2.X, e.Vert2.Y, e.Vert2.Z,
		e.Dist)
}

func solve(data string, n int, part int) (prod int) {
	boxes := make([]*vert, 0)
	edges := make([]edge, 0)
	for _, row := range strings.Split(data, "\n") {
		coords := strings.Split(row, ",")
		box := &vert{X: toFloat(coords[0]), Y: toFloat(coords[1]), Z: toFloat(coords[2])}
		box.Parent = box // for disjoint set
		box.Size = 1     //itself
		for _, b := range boxes {
			d := dist(*box, *b)
			e := edge{Vert1: b, Vert2: box, Dist: d}
			edges = append(edges, e)
		}
		boxes = append(boxes, box)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Dist < edges[j].Dist
	})

	merges := 0
	for _, e := range edges {
		if merges >= n {
			break
		}
		r1 := e.Vert1.GetRoot()
		r2 := e.Vert2.GetRoot()
		if r1 != r2 {
			r1.Parent = r2
			r2.Size += r1.Size
			if r2.Size == len(boxes) && part == 2 {
				return int(e.Vert1.X) * int(e.Vert2.X)
			}
		}
		merges++
	}

	sizes := []int{}
	seen := map[*vert]bool{}
	for _, b := range boxes {
		root := b.GetRoot() // important
		if !seen[root] {
			sizes = append(sizes, root.Size)
			seen[root] = true
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}
