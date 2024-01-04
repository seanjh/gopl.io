package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			fmt.Fprintln(os.Stderr, "opening %s\n", arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, file_counts := range counts {
		total := 0
		files := make([]string, 0, len(files))
		for filename, n := range file_counts {
			total += n
			if n > 0 {
				files = append(files, filename)
			}
		}
		if total > 1 {
			fmt.Printf("%d\t\"%s\"\t(files: %s)\n", total, line, strings.Join(files, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		file_counts, ok := counts[line]
		if !ok {
			file_counts = make(map[string]int)
			counts[line] = file_counts
		}
		counts[line][f.Name()]++
	}
}
