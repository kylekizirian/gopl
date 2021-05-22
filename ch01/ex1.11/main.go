// Try fetchall with longer argument lists, such as samples
// from the top million web sites available at alexa.com.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := getWebsiteUrls(os.Args[1])

	for _, url := range urls {
		go fetch(url, ch) // start goroutine
	}

	fmt.Println("time\tbytes\tdomain")
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Prepends http:// if not already at start of URL
func prependHttp(url string) string {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	return url
}

// Given CSV file with top 500 websites, return string
// slice with URLs
func getWebsiteUrls(csvFile string) []string {
	file, err := os.Open(csvFile)

	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s", csvFile))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse provided CSV file")
	}

	urls := make([]string, len(lines))

	for index, line := range lines {
		urls[index] = prependHttp(line[1])
	}

	return urls
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d\t%s", secs, nbytes, url)
}

// Prints error and exits with code 1
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
