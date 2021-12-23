package concurrency

// Concurrency
// Concurrency means loading more Go Routines at a time. If a routine is blocked, then another one is executed.
// On a single CPU you can run only concurrent application, not parallel execution
// Parallelism means executing multiple Go Routines at the same time and this requires multiple CPUs
// Concurrency => Independently executing processes / dealing with multiple things at once
// Parallelism => Simultaneous execution of processes

// Main function in Concurrency
// In Go, the main function is a Go Routine
// The main function doesn't wait for Go Routines to start or finish their executions
// If the main execution is finished, then the program will exit
// Therefore, it's possible that the routines are not executed yet
// This should be handled with the use of synchronization => WaitGroups

// Checking Current Operation
// runtime.NumCPU() => logical CPUs, including number of cores
// runtime.NumGoroutine() => active Go Routines
// runtime.GOOS and runtime.GOARCH for architecture and OS related info
