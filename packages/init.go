package packages

// Init Function
// Init doesn't have any args or return anything
// Called automatically after a package is initialized and called before main function
// No need to call it explicitly and cannot call it explicitly
// Can declare more than one init functions
// Executed on the order of declarations
// Useful to initialize global variables that can't be initialized in the global context
// Inside the init can process a global variables, e.g. use for loop, since Go doesn't allow expression usage outside function
