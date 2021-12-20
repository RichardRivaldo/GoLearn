package structures

// Scope means visibility
// Lifetime => The interval that an identifier exists as the program executes
// An identifier can't be declared under the same scope, but can be declared in other scopes
// Scopes in Go: File, Package, Block / Local scope

// import in Go is file scoped, must be imported again in another file

// Variables and constant declared outside functions are package-scoped
// Meaning that programs in the package can access it
// Unused variables in package scope won't be an error, block scope will
