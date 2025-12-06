package main

import (
	"aog/internal/aogutils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result := solve1(data)
	result2 := solve2(data)
	fmt.Println(result, result2)
}

func solve1(data string) (sum int) {
	rows := strings.Split(data, "\n")
	problems := make([][]string, len(strings.Fields(rows[0])))
	for _, row := range strings.Split(data, "\n") {
		for x, cell := range strings.Fields(row) {
			problems[x] = append(problems[x], cell)
		}
	}
	return calc_problem(problems)
}

func solve2(data string) (sum int) {
	rows := strings.Split(data, "\n")
	arrangement := make([][]string, len(rows[0]))
	for _, row := range strings.Split(data, "\n") {
		for x := 0; x < len(row); x++ {
			cell := row[x]

			arrangement[x] = append(arrangement[x], string(cell))
		}
	}

	problems := make([][]string, 1)
	cur := 0
	re := regexp.MustCompile(`\s+`)

	for i := len(arrangement) - 1; i >= 0; i-- {
		field := strings.Join(arrangement[i], "")
		field = re.ReplaceAllString(field, "")

		if field == "" {
			cur += 1
			problems = append(problems, make([]string, 0))
			continue
		}
		last := field[len(field)-1]
		if last == '+' || last == '*' {
			problems[cur] = append(problems[cur], field[:len(field)-1])
			problems[cur] = append(problems[cur], string(last))
		} else {
			problems[cur] = append(problems[cur], field)
		}
	}
	return calc_problem(problems)
}

func calc_problem(problems [][]string) (sum int) {
	for _, p := range problems {
		fmt.Println(p)
		if p[len(p)-1] == "+" {
			sub := 0
			for _, num := range p {
				if num == "+" {
					break
				}
				n, _ := strconv.Atoi(num)
				sub += n
			}
			sum += sub
		} else {
			sub := 1
			for _, num := range p {
				if num == "*" {
					break
				}
				n, _ := strconv.Atoi(num)
				sub *= n
			}
			sum += sub
		}
	}
	return
}
