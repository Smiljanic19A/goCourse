package main

import "fmt"

func main() {

	//var colors map[string]string
	colors := make(map[string]string)
	//colors := map[string]string{
	//	"red":   "#ff0000",
	//	"green": "#00ff00",
	//}

	colorsTrings := []string{"red", "blue", "green", "yellow"}

	for _, color := range colorsTrings {
		colors[color] = color
	}

	delete(colors, "red")
	fmt.Println(colors)
}
