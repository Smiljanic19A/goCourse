package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("You must provide exactly 1 filename")
		os.Exit(1)
	}

	filename := args[0]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, er := io.Copy(os.Stdout, file)

	if er != nil {
		fmt.Println(err)
	}
	// fmt.Println(args)
}
