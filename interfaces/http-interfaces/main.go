package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	w, e := io.Copy(os.Stdout, resp.Body)
	if e != nil {
		panic(e)
	}
	fmt.Println(w)

}
