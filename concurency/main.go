package main

import (
	"fmt"
	"net/http"
)

var websites []string = []string{"" +
	"https://google.com",
	"https://facebook.com",
	"https://golang.org",
	"https://instagram.com",
	"https://ugaaaaaaaavyuuiasdaaaaaaa.com",
}
var c chan string = make(chan string)

func main() {
	for _, website := range websites {
		go checkLink(website)
	}
	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string) {
	resp, err := http.Get(link)
	if err != nil {
		errString := "Error occured calling " + link + " Error" + err.Error()
		c <- errString
		return
	}

	if resp.StatusCode == 200 {
		c <- "All good with: " + link
		return
	}

	c <- "Anomaly detected for " + link + " Non 200 status code received" + resp.Status
}
