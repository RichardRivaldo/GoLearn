package basic

// Function
// camelCase
// Function name in same package must be unique
// No overloading
// Predefined functions: main() and int()
// Support first class function, that is assign a function to a variable or returns a function

// Declaration
// func name(parameters) returns {}

// Parameters, Args, and Returns
// Passed by value (copy), no default parameters in Go
// Error value should be the last return value
// Named return value will be zero-valued at start
// Naked return => return without specifying return values, should only be used in simple functions
// func name(param1 type1, param2 type2) returnType
// func name(param1, param2 sameType) (returnType1, returnType2)
// func name(params) (namedReturn returnType)

// Variadic Functions
// Variable number of arguments, zero or more
// Only last parameter of a function should be variadic
// Zero arguments invocation will cause the parameter to be nil
// func name(param ...type)

// Anonymous Function
// Doesn't have any name and declared inline, IIEE Function like JavaScript
// func(params){}(args)
