package main

import (
	"os"
)

func main() {
	arguments := os.Args[1:]
	URL, path, count := parseInput(arguments)
	t, b, e := handleRequest(URL, path, count)
	if count != 0 {
		printOutput(t, b, e)
	}
}
