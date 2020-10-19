package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
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
		return val.Hostname(), val.Path, false
	}
	return "", "", true
}

func urlError() {
	fmt.Println("Enter valid URL")
}

func sendRequest(url string, path string) {
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
	// Check if any errors occured
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}

	// fmt.Println(responseString)

	fmt.Println(len(responseString))

}

func handleRequest(url string, path string, count int) {
	timetaken := make([]int, 0, count)
	// bytes := make([]int, 0, count)
	// errors := make([]int, 0, count)

	for i := 0; i < count; i++ {
		start := time.Now()
		sendRequest(url, path)
		end := time.Since(start)

		timetaken = append(timetaken, int(end.Milliseconds()))
	}

}

func main() {
	arguments := os.Args[1:]
	URL := ""
	count := 1
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
			fmt.Print(c)
		}
	}

	// fmt.Print(URL+" "+path, count)

	handleRequest(URL, path, count)
	total := `HTTP/1.1 200 OK
	Date: Mon, 19 Oct 2020 00:01:27 GMT
	Content-Type: application/json
	Content-Length: 190
	Connection: close
	Set-Cookie: __cfduid=dccbdf1937992b8cb6dbcd71bb8e5d33e1603065687; expires=Wed, 18-Nov-20 00:01:27 GMT; path=/; domain=.endeavor.workers.dev; HttpOnly; SameSite=Lax
	cf-request-id: 05dfc2465600002ac0dc83a000000001
	Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
	Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report?lkg-colo=16&lkg-time=1603065688"}],"group":"cf-nel","max_age":604800}
	NEL: {"report_to":"cf-nel","max_age":604800}
	Server: cloudflare
	CF-RAY: 5e463983b9e02ac0-IAD`
	s := `[{"name":"Youtube","url":"https://www.youtube.com"},{"name":"Cloud Flare","url":"https://www.cloudflare.com"},{"name":"Among US","url":"https://store.steampowered.com/app/945360/Among_Us/"}]`
	print("len of something ", len(s))
	print("len of total", len(total))
}
