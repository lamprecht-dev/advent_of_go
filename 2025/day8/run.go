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
	result := solve1(data, 1000)
	result2 := solve2(data)
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
	return v
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

func makeEdge(v1, v2 *vert) edge {
	dx, dy, dz := v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z
	return edge{
		Vert1: v1,
		Vert2: v2,
		Dist:  math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2) + math.Pow(dz, 2)),
	}
}

func solve1(data string, n int) (prod int) {
	boxes := make([]vert, 0)
	edges := make([]edge, 0)
	for _, row := range strings.Split(data, "\n") {
		coords := strings.Split(row, ",")
		box := vert{X: toFloat(coords[0]), Y: toFloat(coords[1]), Z: toFloat(coords[2])}
		box.Parent = &box // for disjoint set
		box.Size = 1      //itself
		for _, b := range boxes {
			e := makeEdge(&box, &b)
			edges = append(edges, e)
		}
		boxes = append(boxes, box)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Dist < edges[j].Dist
	})

	found := 0
	for _, e := range edges {
		if found >= n {
			break
		}
		r1 := e.Vert1.GetRoot()
		r2 := e.Vert2.GetRoot()
		if r1 != r2 {
			found++
			r1.Parent = r2
			r2.Size += r1.Size
		}
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
	prod = sizes[0] * sizes[1] * sizes[2]

	return
}

func solve2(data string) (sum int) {
	return
}
