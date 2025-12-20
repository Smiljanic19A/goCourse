package main

import "fmt"

func main() {
	sliceOfInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range sliceOfInts {
		if i%2 == 0 {
			fmt.Println("Even")
			continue
		}
		fmt.Println("Odd")
	}
}
