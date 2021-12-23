package concurrency

// Go Routines
// goroutine is a lightweight thread of execution
// goroutine is also function that is capable of running concurrently with other functions
// goroutine is very cheap and smaller than usual threads, since it takes only around 2KB of stack space, while threads take 1-2MB space.
// goroutine has its own execution stack that shrinks and grows as needed, not like fixed-size OS Thread stack
// Scheduling is cheaper with goroutine, it uses Go scheduler with m:n scheduling that multiplexes m goroutines on n OS threads
// goroutine has no identity
// go functionInvocation() => Spawning a goroutine
