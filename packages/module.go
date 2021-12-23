package packages

// Modules
// Module is a collection of related Go packages stored in a directory tree with a go.mod file at its root
// Module contains one or more packages
// go.mod defines the module path, import path, and dependencies
// go command has supports to work with modules, including resolving dependencies and having multiple versions

// Initializing Modules
// go.mod is basically package.json in NodeJS
// go mod init directory

// Importing Modules
// Modules path is usually version controls like Github
// Import the package path, not the module path if there are more than 1 packages inside a module
// import "pathWithoutHTTPS"

// Building Modules
// Building modules mean downloading the modules hosted on the VCS to GOPATH directory,
// also importing the packages to be used by the program with require command inside go.mod
// This will also generate and verify checksum for the packages in go.sum file
// IMPORTANT: go run main.go usual command will also do the same
// go build

// Creating Go Modules
// 1. Create a repository
// 2. Create the packages that want to be published inside the module
// 3. Can separate the direcotry of the packages if needed
// 4. Initialize the module by executing `go mod init githubRepoURL`

// Publishing Go Modules
// Simply push the packages to the repository
// Create a release with semantic versioning

// Semantic Versioning
// Major version => Backward incompatible changes
// Minor version => Backward compatible changes
// Patch version => Bug fixes for backward compatible version
// vMAJOR.MINOR.PATCH

// Dependencies Versioning
// Can directly change the version inside go.mod and build the package
// Can also run go get commands
// go get -u => Minor Version Update
// go get -u githubModulePath@version => Major and Patch Version Update

// Multi-Dependency Package
// Possible to use multi versions package
// Can create alias for the imported package
// import (
// 	v1Alias "v1PackagePath"
// 	v2Alias "v2PackagePath"
// )
