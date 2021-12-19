package variables

// Go Variables
// Type inference in Go, so we don't need to specify type in declaration
// Go only knows the type of the variable when compiled
// Variable values will be taken care of in runtime processes
// Strong typed language, can assign float to int variable

// Declaration
// Capital case for first letter of a variable indicates that
// it will be exported and can be used by other packages
// keyword identifier type (opt. = value)
// var variableName int = 9

// Blank Identifier
// All variables must be used in Go
// _ can be used to assign unused variables
// _ = unusedVariable

// Short declaration
// variable_name := 9
// Can be used to declare new variable
// Can't be used for package-scoped declaration
// To change an already declared variable, just assign it with = operator
// declaredVariable = newValue

// Multiple Declaration
// Short Way
// var1, var2 := val1, val2
// At least one variable should be a new variable

// Normal Way
// var var1, var2, var3 type. var1 var2 var3 all have same type
// var (var1 type var2 type var3 type)

// Multiple Assignment
// Tuple-like assignment, can swap two variables' values like Python
// var1, var2 = val1, val2 => var1, var2 = val2, val1
