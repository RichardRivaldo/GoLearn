package basic

// Types in Go
// Statically and strongly typed language
// Can change variable type once declared
// Variable in Go is always zero-valued
// pointer => nil, string = "", int/float = 0

// Convention
// Camel Case
// Use short naming variables
// Start a variable with _ or letters

// Type bool, true or false
// String => unicode chars, ""
// Rune => ''

// Numeric Values
// int for signed, uint for unsigned, consists of 8, 16, 32, and 64 bits
// float32 or float64, can omit zero near the . (ex: .-2, 2.)
// complex64 and complex128
// Can use _ for large numbers

// Composites

// Array
// Same type elements, fixed length when declaration
// var arr = [n]type{values}

// Slice
// Same type, variable length
// var arr = []type{values}

// Map
// Key-Value
// map := map[keyType]valType{Key: Values}

// Struct
// Sequence of named elements / fields
// type typeName struct {elements types}

// Pointer
// Stores memory addr of a variable, uninitialized value is nil
// ptr := &anotherVariable

// Function, Interface, Channel

// Conversions (not casting, returns new value)
// Cannot operate two operands with diff name even if the type is the same (ex: int * int64)
// string(int) will return char based on the ASCII value, conversion with float doesn't work
// Sprintf from fmt package is useful for type conversions from int or float to string
// String to int or float can also use stringconv.parse(Float, Bool, Int)
// There is also AToI and IToA (ASCII to Int and reversed)

// Named / Defined Types
// Use underlying or source type, that is the predefined type
// Not an alias, entirely new type and must be converted to be operated with predefined types
// Can attach methods to the defined type
// There is no-type hierarchy in Go
// type namedType underlyingType

// Aliases
// The identifier will be the same type with the source type
// Can be used without conversions
// int and uint are default aliases for 32 or 64 bits, depend on the platform
// byte (uint8), rune (int32) => distinguish chars and numbers
// type aliasIdentifier = sourceType
