package main

import "fmt"

func main() {
	t := triangle{
		height: 4.0,
		base:   2.0,
	}

	s := square{
		sideLength: 5.0,
	}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
