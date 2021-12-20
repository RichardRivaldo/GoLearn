package datstruct

// Array
// Composite and indexable elements
// Fixed length, identical element types
// Can modify the element by index with simple assignment
// Elements in array are stored contiguously in Go
// The length is allocated on compile time

// Declaration
// Always zero-valued if not initialized
// Not mandatory to init all elements
// var arrName [length]type
// var arrName = [length]type{val1, val2, val3}
// short := [length]type{val1, val2, val3}
// Multiple line declaration must use , at the last element of the array

// Ellipsis Operator
// Can not specify any length and Go will count the length based on the elements
// short := [...]{val1, val2, val3, ...}

// Keyed Declaration
// Specify index upon declaration for each array elements
// The elements will be automatically ordered based on the index key (must be positive integers)
// Element which is not specified will get its index from the last indexed element declared before it
// Unspecified elements will also be zero-valued
// arr := [...]type{key2: val1, key1: val2}

// Iteration
// Normal index loop
// for index, element := range arr {}

// Operators
// == for equality => must be same length, same elements, same order
// := => Copy assignment, copy the target array

// Operations
// Length => len(arr)
