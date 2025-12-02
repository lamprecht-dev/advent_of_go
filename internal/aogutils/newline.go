package aogutils

import (
	"bufio"
	"strings"
)

type LineScanner struct {
	scanner *bufio.Scanner
}

func (l LineScanner) Scan() bool {
	return l.scanner.Scan()
}

func (l LineScanner) Text() string {
	return l.scanner.Text()
}

func NewLines(data string) LineScanner {
	return LineScanner{bufio.NewScanner(strings.NewReader(data))}
}
