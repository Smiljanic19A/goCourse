package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logger struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	w, e := io.Copy(logger{}, resp.Body)
	if e != nil {
		panic(e)
	}
	fmt.Println(w)

}

func (l logger) Write(bs []byte) (int, error) {
	//fmt.Println(string(bs))
	os.WriteFile("log.txt", bs, 777)
	return len(bs), nil
}
