package main

import (
	"aog/internal/aogutils"
	"fmt"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result := solve1(data)
	result2 := solve2(data)
	fmt.Println(result, result2)
}

type PathCount struct {
	None int
	All  int
	Dac  int
	Fft  int
}

type Paths map[string]map[string]bool

func parse(data string) Paths {
	paths := make(Paths)
	for _, row := range strings.Split(data, "\n") {
		parts := strings.Split(row, ":")
		key := parts[0]
		if _, ok := paths[key]; !ok {
			paths[key] = make(map[string]bool)
		}
		for _, target := range strings.Fields(parts[1]) {
			targetNode := target
			paths[key][targetNode] = true
		}
	}
	return paths
}

func solve1(data string) int {
	paths := parse(data)

	out := CalcPaths("you", paths, make(map[string]PathCount))
	return out.All + out.None + out.Fft + out.Dac
}

func solve2(data string) int {
	paths := parse(data)

	out := CalcPaths("svr", paths, make(map[string]PathCount))
	return out.All
}

func CalcPaths(cur string, paths Paths, cache map[string]PathCount) PathCount {
	if cur == "out" {
		return GetPathCount(1, 0, 0, 0)
	}
	if pc, ok := cache[cur]; ok {
		return pc
	}

	sum := GetPathCount(0, 0, 0, 0)
	for t := range paths[cur] {
		tpaths := CalcPaths(t, paths, cache)
		switch t {
		case "dac":
			sum.All += tpaths.All + tpaths.Fft
			sum.Dac += tpaths.Dac + tpaths.None
		case "fft":
			sum.All += tpaths.All + tpaths.Dac
			sum.Fft += tpaths.Fft + tpaths.None
		default:
			sum.All += tpaths.All
			sum.None += tpaths.None
			sum.Fft += tpaths.Fft
			sum.Dac += tpaths.Dac
		}
	}

	cache[cur] = sum

	return sum
}

func GetPathCount(none, fft, dac, all int) PathCount {
	return PathCount{All: all, None: none, Fft: fft, Dac: dac}
}
