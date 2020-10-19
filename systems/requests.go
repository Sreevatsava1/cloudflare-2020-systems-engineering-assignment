package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func getResponse(url string, path string) string {
	timeout, _ := time.ParseDuration("10s")

	dialer := net.Dialer{
		Timeout: timeout,
	}

	conn, err := tls.DialWithDialer(&dialer, "tcp", url+":https", nil)

	if err != nil {
		fmt.Println(err.Error())
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
	errors := make([]int, 0)

	if count == 0 {
		response := getResponse(url, path)
		println("----------------Response----------------")
		fmt.Println(response)
	}

	for i := 0; i < count; i++ {
		start := time.Now()
		response := getResponse(url, path)
		end := time.Since(start)

		times = append(times, int(end.Milliseconds()))
		b, e := parseResponse(response)
		bytes = append(bytes, b)
		if e != 200 {
			errors = append(errors, e)
		}

	}

	return times, bytes, errors

}
