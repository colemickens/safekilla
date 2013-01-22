package main

import (
	//"errors"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

//var agent *string = flag.String("agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.22 (KHTML, like Gecko) Chrome/19.0.1049.3 Safari/535.22", "User-Agent to emulate")
var agent *string = flag.String("agent", "wii", "User-Agent to emulate")
var continuous *int64 = flag.Int64("c", -1, "Run every N minutes (-1 means don't repeat)")
var url *string = flag.String("url", "http://www.google.com", "URL to ping")

func setHeader(req *http.Request) {
	req.Header.Set("User-Agent", *agent)
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept", "*/*")
}

func Ping() {
	blocked := true

	client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirected")
			blocked = true
			setHeader(req)
			return nil
		},
	}

	for blocked == true {
		fmt.Println("http get", *url)

		req, _ := http.NewRequest("GET", *url, nil)
		setHeader(req)

		blocked = false
		_, err := client.Do(req)

		if err != nil {
			fmt.Println("The attempt to ping failed. Check your internet connection.")
		}

		time.Sleep(1e9)
		fmt.Println()
	}
	fmt.Println("[" + *url + "] is directly reachable")
}

func main() {
	fmt.Println("SafeConnect Killa")
	flag.Parse()

	Ping()

	if *continuous > 0 {
		for {
			fmt.Println("\nsleeping for", *continuous, "minutes")

			d := 1e9 * 60 * (*continuous)
			time.Sleep(time.Duration(d))

			Ping()
		}
	}
}

// wget

// GET / HTTP/1.0\r\n
// User-Agent: wii\r\n
// Accept: */*\r\n
// Host: www.yahoo.com\r\n
// Connection: Keep-Alive\r\n
// \r\n

// safekilla

// GET / HTTP/1.1\r\n
// Host: www.google.com\r\n
// User-Agent: wii\r\n
// Accept: */*\r\n
// Connection: Keep-Alive\r\n
// Accept-Encoding: gzip\r\n
// \r\n
