package oop

// Method
// Named type can have methods or also called receiver function
// Go doesn't have this keyword, instead uses the receiver as this replacement
// Method in Go is function that receives the receiver as the function argument
// Method belongs to the type while function belongs to the package

// Creating
// Basically creating a function with the type as receiver
// func (receiver type) funcName(args) returnType{}

// Invoking
// object.method()
// type.method(receiver)

// Pointer Receiver
// Pointer receiver is valid except for pointer types / aliases
// Setter should use pointer receiver as normal receiver is also passed by value
// (&object).method()
// object.method() => The object is implicitly converted to pointers
// (*object).method() => The object itself is a pointer, basically dereferencing the address and call the method
