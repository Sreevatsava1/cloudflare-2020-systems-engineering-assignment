package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"sort"
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
		return val.Hostname(), val.Path, false
	}
	return "", "", true
}

func getResponse(url string, path string) string {
	timeout, _ := time.ParseDuration("10s")

	dialer := net.Dialer{
		Timeout: timeout,
	}

	conn, err := tls.DialWithDialer(&dialer, "tcp", url+":https", nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}

	defer conn.Close()

	conn.Write([]byte("GET " + path + " HTTP/1.0\r\nHost: " + url + "\r\n\r\n"))

	response, err := ioutil.ReadAll(conn)
	responseString := string(response)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}

	return responseString

}

func parseResponse(response string) (int, int) {
	b := len(response)
	r := strings.Split(response, " ")
	status, _ := strconv.Atoi(r[1])
	return b, status
}

func handleRequest(url string, path string, count int) ([]int, []int, []int) {
	times := make([]int, 0, count)
	bytes := make([]int, 0, count)
	errors := make([]int, 0, count)

	if count == 0 {
		response := getResponse(url, path)
		print(response)
	}
	for i := 0; i < count; i++ {
		start := time.Now()
		response := getResponse(url, path)
		end := time.Since(start)

		times = append(times, int(end.Milliseconds()))
		b, e := parseResponse(response)
		bytes = append(bytes, b)
		errors = append(errors, e)
	}

	return times, bytes, errors

}

func parseInput(arguments []string) (string, string, int) {
	return "", "", 0
}

func printOutput(times []int, bytes []int, errors []int) {
	count := len(times)
	sort.Ints(times)
	sort.Ints(bytes)
	tott := 0
	totb := 0
	for i := 0; i < count; i++ {
		tott += times[i]
		totb += bytes[i]
	}
	println("Number of requests : ", count)
	println("Fastest time (ms) : ", times[0])
	println("Slowest time (ms) : ", times[count-1])
	println("Mean time (ms) : ", (tott*1.0)/count)
	println("Median time (ms) : ", times[count/2])
	println("Percentage of successful requests : ", ((count - len(errors)*100.0) / count))
	println("Median time (ms) : ", times[count/2])
	println("Error codes : ", errors)
	println("Size of bytes of smallest response : ", bytes[0])
	println("Size of bytes of largest response : ", bytes[count-1])
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
	if arguments[2] == "--profile" {
		if len(arguments) > 3 {
			c, err := strconv.Atoi(arguments[3])
			if err != nil && c > 0 {
				fmt.Println("Enter valid request count")
			}
			count = c
			fmt.Print(c)
		}
	}
	if count != 0 {
		t, b, e := handleRequest(URL, path, count)
		printOutput(t, b, e)
	}
}
