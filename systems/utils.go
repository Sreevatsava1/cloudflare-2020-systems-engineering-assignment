package main

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
)

func urlError() {
	/*
		Prints error message
	*/
	fmt.Println("Enter valid URL")
}

func parseURL(URL string) (string, string, bool) {
	/*
		Parses the URL and returs the hostname, pathname and isURLValid
	*/
	URL = strings.ToLower(URL)
	if URL[:4] != "http" {
		URL = "https://" + URL
	}
	val, err := url.Parse(URL) // Parses the URL using url function
	if err == nil && (val.Scheme == "http" || val.Scheme == "https") && val.Host != "" {
		return val.Hostname(), val.Path, false
	}
	return "", "", true
}

func helpInstructions() {
	/*
		Prints help instructions
	*/
	println("-------------------Usage-------------------\n\nTo get HTTP response from a URL\nCode : go run . --url <URL>")
	println("\nExample:\ngo run . --url https://www.google.com")
	println("\nIf you need to profile the url\nCode : go run . --url <URL> --profile <no. of requests>")
	println("\nExample:\ngo run . --url https://www.google.com --profile 10")
}

func parseInput(args []string) (string, string, int) {
	/*
		Parses the input arguments from the command line and returns the URL, pathname and no. of requests count
	*/

	//Initializing values
	URL := ""
	count := 0
	path := "/"

	if len(args) == 0 || args[0] == "--help" {
		helpInstructions()
		os.Exit(1)
	}

	for i := 0; i < len(args); i += 2 {
		if args[i] == "--url" {
			if i+1 < len(args) {
				URL = args[i+1]
				u, p, err := parseURL(URL)
				if err == true {
					urlError()
					os.Exit(1)
				}
				URL = u
				if p != "" {
					path = p
				}
			} else {
				println("Enter an url after the --url")
				println()
				helpInstructions()
				os.Exit(1)
			}

		} else if args[i] == "--profile" {
			c, err := strconv.Atoi(args[i+1])
			if err != nil || c <= 0 {
				fmt.Println("Enter valid request count")
				println()
				helpInstructions()
				os.Exit(1)
			}
			count = c
		} else {
			println("Your input is invalid, follow the instructions")
			println()
			helpInstructions()
			os.Exit(1)
		}

	}
	return URL, path, count
}

func printOutput(times []int, bytes []int, errors []int) {
	/*
		Prints the final output profile for each website
	*/
	count := len(times)
	sort.Ints(times)
	sort.Ints(bytes)
	tott := 0
	totb := 0
	for i := 0; i < count; i++ {
		tott += times[i]
		totb += bytes[i]
	}
	fmt.Println("\n-----------------------Profile-----------------------")
	println()
	fmt.Println("Number of requests : ", count)
	fmt.Println("Fastest time (ms) : ", times[0])
	fmt.Println("Slowest time (ms) : ", times[count-1])
	fmt.Println("Mean time (ms) : ", (float64(tott)*1.0)/float64(count))
	fmt.Println("Median time (ms) : ", times[count/2])
	fmt.Println("Percentage of successful requests : ", (((float64(count) - float64(len(errors))) * 100.0) / float64(count)), "%")
	fmt.Println("Median time (ms) : ", times[count/2])
	if len(errors) > 0 {
		fmt.Println("Error codes : ", errors)
	} else {
		fmt.Println("Error codes : ")
	}
	fmt.Println("Size of bytes of smallest response : ", bytes[count-1])
	fmt.Println("Size of bytes of largest response : ", bytes[0])
	println()
	println()
}
