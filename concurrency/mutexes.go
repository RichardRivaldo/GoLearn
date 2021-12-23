package concurrency

// Race Condition
// Accessing an address by two or more routines at the same time
// Handled with Mutex or Channels

// Race Detector
// Race Condition detectors built based on C/C++
// go run -race main.go

// Mutexes (Mutual Exclusion)
// Explicit Synchronization, prevent race conditions
// var mutex sync.Mutex
// m.Lock() ... m.Unlock() => The code between them will only be executed by one routine
// Can also use defer m.Unlock() after m.Lock()
