package basic

// Conditional
// Keywords: if, else if, else
// No truthy values in Go, can only use boolean
// No need parentheses

// Simple Statement
// Usually used for error handling
// Initialization statement; if statement
// if a := b; a == b {}

// Switch
// Automatically convert the code to if statement
// No need to use break explicitly in Go
// switch var {cases}
// switch true {conditionals}
// switch {conditionals} is the same as switch true {conditionals}
// Can also use simple statements like if

// Loop
// Only for, no while
// Syntax is like C
// for init; limit; steps {}
// The step can be omitted or put inside the loop code
// Can handle multiple variables in the statement
// Can use continue and break normally

// while
// Basic While => for condition {changeIterHere}
// While True => for {}

// Labels
// Identifier for loops, must be used, can be used in break, continue, or goto statement
// Doesn't conflict with other identifiers like variable or function names
// Usually used to terminate outer loops from inner one
// Usage: labelName : forloop{ innerloop {break labelName}}

// goto
// Can be used to specify which label the program want to execute
// Not limited to loops and switch
// Illegal to jump over variable declarations after a goto statement
