package oop

// Interface
// Same concept, a contract every type that extends the interfaces should implement
// Allows polymorphism
// In Go, if a type implements all the method inside an interface, it automatically implements the interface
// A type can implement one or more interfaces, and is a type of the interface if it implements it
// Zero-valued into nil and has abstract values
// Cannot implement or extend another interfaces, but can be merged with interface embedding

// Declaration
// type name interface {functionName(params) returnType}

// Polymorphism
// Interface is dynamic type that can change during runtime, and makes it possible to implement polymorphism
// var i interfaceName => nil
// i = anyObjectFromAnyTypeThatImplementsTheInterface

// Type Assertion
// If a type that implements the interface has another method, then an interface object cannot invoke that method (like TypeScript)
// Type assertion is the answer to that. Use the Comma OK Idiom
// typeObject, assertionOK := interfaceObject.(typeImplementTheInterface)

// Type Switches
// Basically switches or if based on the type of the interface
// Using the type keyword for type switches
// switch val := interfaceObject.(type) {case typeImplementTheInterface}

// Interface Embedding
// type i interface {interface1, interface2, otherMethods()}
// Circular embedding is not allowed!

// Empty Interface
// Interface with no method, all type implements the interface and is valid for that interface
// Will still need type assertion
// Useful for unknown arguments type or all types are supposed to be allowed
// Hard to maintain as usual :D
// interface{} => empty interface type
// type i interface{} => named empty interface
