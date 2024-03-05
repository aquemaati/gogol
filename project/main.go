package main

/*
"Using a main.go file in a Go project serves as the designated entry point for executable applications,
adhering to a widely recognized convention, facilitating clarity, organization, and executable output generation."
*/
import (
	"fmt"
	mylibrary "project/pkg/mylibrary"
)

func main() {
	greet := mylibrary.Greet("project")
	fmt.Println(greet)
}