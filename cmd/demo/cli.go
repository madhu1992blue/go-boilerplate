package main

import "fmt"

// This will be set with -ldflags="-X main.version=value" at the build time.
// It need not be exported. Hence, keeping in lowercase.
var version string

// main function in main package.
// filename doesn't matter.
func main() {
	fmt.Println("Version:", version)
}
