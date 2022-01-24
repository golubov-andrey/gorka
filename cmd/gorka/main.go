// Copyright (c) 2022 Golubov Andrey
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

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
