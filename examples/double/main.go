package main

import "github.com/David-Rushton/pretty"

func main() {
	// Demonstrating the use of a double width and height formats.

	pretty.Printf("Double hight line\n", pretty.WithDoubleHeightTop())
	pretty.Printf("Double hight line\n", pretty.WithDoubleHeightBottom())
}
