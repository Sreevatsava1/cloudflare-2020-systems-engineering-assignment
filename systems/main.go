package main

import (
	"os"
)

func main() {
	/*
	 Main program
	*/
	arguments := os.Args[1:]
	URL, path, count := parseInput(arguments)               // Gets url, pathname and count from the parseInput funcion that takes command line arguments as parameters
	times, bytes, errors := handleRequest(URL, path, count) // Gets timetaken, bytes and errors arrays from handleRequest
	if count != 0 {
		printOutput(times, bytes, errors) // Prints the Profile of the given website
	}
}
