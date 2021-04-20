// Modify dup2 to print the names of all files in which each duplicated
// line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Represents a set with methods `add` and `has`
type filesSet map[string]bool

// Returns whether string `file` is in filesSet
func (f filesSet) String() string {
	files := []string{}

	for file, _ := range f {
		files = append(files, file)
	}

	return strings.Join(files, " ")
}

// Adds a file name to the filesSet
func (f filesSet) add(file string) {
	f[file] = true
}

// Struct representing number of times a line has been
// seen and what files it has been seen in
type lineInfo struct {
	Count     int
	Filenames filesSet
}

// Returns a new lineInfo object with the count incremented
// and the filename added to the set of files it's in
func (l lineInfo) update(filename string) lineInfo {
	updated := l
	updated.Count++
	updated.Filenames.add(filename)
	return updated
}

// Maps lines to the line's counter and set of files that line
// is contained in
type lineInfoMap map[string]lineInfo

// Returns whether line is already in the info map
func (l lineInfoMap) has(line string) bool {
	_, ok := l[line]
	return ok
}

func main() {
	counts := make(lineInfoMap)
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, arg)
		f.Close()
	}

	// For each duplicate line (count is greater than 1), print the number of
	// times it was seen, then line itself, and the files that contain the line
	for line, info := range counts {
		if info.Count > 1 {
			fmt.Printf("%d\t%s\t%s\n", info.Count, line, info.Filenames)
		}
	}
}

func countLines(f *os.File, counts lineInfoMap, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()

		// Initialize entry in map if this is the first time we've
		// seen this line
		if !counts.has(line) {
			counts[line] = lineInfo{
				Count:     0,
				Filenames: filesSet{filename: true},
			}
		}

		// Increment the line's counter and add this file to
		// to the set of files that contain this line
		counts[line] = counts[line].update(filename)
	}
}
