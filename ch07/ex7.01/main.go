package main

import (
	"bufio"
	"bytes"
)

type WordCount int
type LineCount int

func scanCount(p []byte, split bufio.SplitFunc) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(split)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func (w *WordCount) Write(p []byte) (int, error) {
	n, err := scanCount(p, bufio.ScanWords)
	if err != nil {
		return 0, err
	}
	*w += WordCount(n)
	return n, nil
}

func (l *LineCount) Write(p []byte) (int, error) {
	n, err := scanCount(p, bufio.ScanLines)
	if err != nil {
		return 0, err
	}
	*l += LineCount(n)
	return n, nil
}
