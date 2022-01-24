package main

import (
	_ "embed"
	"fmt"
)

//go:embed VERSION
var version string

func main() {
	fmt.Println(version)
}
