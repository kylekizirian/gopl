package main

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	var wc WordCount
	n, err := wc.Write([]byte("hello, world!"))
	ok(t, err)
	equals(t, 2, n)
	equals(t, 2, int(wc))

	n, err = wc.Write([]byte("this is my message"))
	ok(t, err)
	equals(t, 4, n)
	equals(t, 6, int(wc))
}

func ok(t *testing.T, err error) {
	if err != nil {
		t.Fatal("err: ", err)
	}
}

func equals(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Fatal("expected: ", expected, " got: ", actual)
	}
}
