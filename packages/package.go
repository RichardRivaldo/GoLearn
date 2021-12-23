package packages

// Go Packages
// Package is basically project directory containing Go files
// All files in a directory must be in the same package
// Avoid _ or camelCase naming, use simple name for packages
// By default, import path is relative to GOPATH/src
// IMPORTANT: Only function name with Capital first case letter can be accessed from other packages

// Executable Package
// Generate executables and can be run
// Predefined name => main package

// Non-Executable Package (libraries or dependencies)
// Used by other package can not be executed (only imported)
// Can have any name
