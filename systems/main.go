package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func helpInstructions() {
	fmt.Println("This is help Message!!")
}

func parseURL(URL string) (string, string, bool) {
	URL = strings.ToLower(URL)
	val, err := url.Parse(URL)
	if err == nil && (val.Scheme == "http" || val.Scheme == "https") && val.Host != "" {
		return val.Hostname(), val.Path, true
	}
	return "", "", false
}

func urlError() {
	fmt.Println("Enter valid URL")
}

func sendRequest(url string) {
	timeout, _ := time.ParseDuration("10s")

	dialer := net.Dialer{
		Timeout: timeout,
	}

	connection, err := tls.DialWithDialer(&dialer, "tcp", url+":https", nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}

	defer connection.Close()
}

func handleRequest(url string, count int) {
	timetaken := make([]int, 0, count)
	// bytes := make([]int, 0, count)
	// errors := make([]int, 0, count)

	for i := 0; i < count; i++ {
		start := time.Now()
		sendRequest(url)
		end := time.Since(start)

		timetaken = append(timetaken, int(end.Milliseconds()))
	}

}

func main() {
	arguments := os.Args[1:]
	URL := ""
	count := 0
	path := ""
	// Check if no arguments, or help flag passed
	if len(arguments) <= 0 || arguments[0] == "--help" {
		helpInstructions()
	}

	// Check if we got 2 arguments, and the 1st is the url flag
	if arguments[0] == "--url" {
		if len(arguments) > 1 {
			URL = arguments[1]
			u, p, err := parseURL(URL)
			if err == true {
				urlError()
				os.Exit(1)
			}
			URL = u
			path = p
		} else {
			urlError()
			os.Exit(1)
		}
	}
	if arguments[2] == "profile" {
		if len(arguments) > 3 {
			c, err := strconv.Atoi(arguments[3])
			if err != nil && c > 0 {
				fmt.Println("Enter valid request count")
			}
			count = c
		}
	}

	handleRequest(URL+path, count)

}
