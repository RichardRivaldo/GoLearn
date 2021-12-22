package datstruct

// Pointer
// Variable that stores memory address of another variable, or nil if it has not been initialized
// No arithmetic like C, but still can read and update the data in the pointer address
// Same & and * usage like C
// Zero-valued to nil

// Declaration
// pointer := &pointedVariable
// var pointer *type => %p for address
// p := new(type) -> p = &x => var x type

// Dereferencing
// *pointer = val

// Pointer to Pointer
// pointer2 := &pointer1
// **pointer2 to dereference pointer2

// Function and Pointer
// Remember: pass by value in Go
// Can only change basic types, array, and struct through functions by using pointer
// For struct, can change directly without using dereference operator as long as the struct object is passed as pointers
// For slices and maps, can change directly since the copy will also reference the same address.
// Don't use pointer for slices and maps!
// Arrays are not meant to be passed to functions, pass slices instead!

// Operator
// == => Equality, true if both are nil or point to same address
