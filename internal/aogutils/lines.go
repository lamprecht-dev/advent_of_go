package aogutils

func Lines(data string) []string {
	lines := []string{}
	ll := NewLines(data)
	for ll.Scan() {
		l := ll.Text()
		lines = append(lines, l)
	}

	return lines
}
