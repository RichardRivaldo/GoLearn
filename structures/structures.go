// Always define packages for each Go files
// Main package is predefined to generate executable standalone Go program
package main

// Import Package
// fmt (r: "fumt") is Go Standard Library

import "fmt"

// Must define function called `main`
// No arguments and no return value
func main() {
	fmt.Println("Hello world")
}

// Compilation
// go run main_file.go
// Only compiles and run the file without any executable

// go build
// Build the current directory into an executable
// Compile file, packages, and dependencies
// Default name is the directory name, specify with go build -o name

// GOOS={windows/linux/darwin} GOARCH=amd64 go build -o name
// Build executable for specific OS and CPU arch

// go install relative_path_to_GOPATH/src
// Place the compiled executable in GOPATH/bin directory

// Formatting
// go fmt -w file.go
// To format a directory, go to parent directory and run gofmt -w -l directory
// or go fmt in the directory
