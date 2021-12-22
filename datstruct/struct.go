package datstruct

// Struct
// Sequence of named elements (fields) or collection of properties
// Go uses struct as complex data structure since it doesnt have class
// The blueprint that defines the struct is fixed at compile time, can't change any fields at runtime
// Can instantiate the struct into an `object`

// Anonymous Struct
// Immediately instantiate the struct without using any aliases
// var := struct{fields}{attributes}

// Anonymous Fields
// Can be mixed with named fields
// Get and set by accessing the type name
// type alias {type1 type2 type3}

// Embedded Structs
// Basically nested structs :)

// Declaration
// Unique fields
// type alias struct {field type}
// Same field type => {field1, field2 type}

// Instantiation
// Uninitialized field will be zero-valued
// alias{attributes} => order matters
// alias{field2: attribute, field1: attribute}

// Get-Set
// object.field
// object.field = val

// Operators
// == => Can compare structs, will compare each corresponding fields
// = => Copy assignment a new struct, will allocate brand new struct in the memory
