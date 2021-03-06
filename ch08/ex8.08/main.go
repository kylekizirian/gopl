package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	inputText := make(chan string)
	eof := make(chan bool)

	go func() {
		for input.Scan() {
			inputText <- input.Text()
		}
		eof <- true
	}()

	for {
		select {
		case text := <-inputText:
			go func() {
				echo(c, text, 1*time.Second)
			}()
		case <-time.After(10 * time.Second):
			return
		case <-eof:
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
