package concurrency

// Channels
// Channels allow routines to communicate, data sent or received must be the same with the declared type
// Zero-valued to nil
// A channel is like pointers that stores an address, passing channels to functions = passing pointers to function
// There are bi- and uni- directional channel. Bi-directional channel can send and receive, while the uni-directional one can only do one operation
// Can convert bidirectional into unidirectional, but not vice versa

// Declaration
// var channel chan type => Not initialized = nil
// channel = make(chan type) => Initialize the channels
// channel := make(chan type)
// make(<- chan type) | make(chan <- type) => Receive-Only | Send-Only
// defer close(channel)

// Operations
// channels <- => sending a value to the channel
// var := <- channel => receive a value from the channel

// GoRoutines and Channels
// go routine(params, channel) => spawn the routine to use channel
// The main function will sleep and wait until a message is sent to the channel
// This is done because the receive operation is a blocking operation
// If there are more than one, then it will wait for messages from each goroutines
// Anonymous function is typically used to spawn the goroutine. Remember to use blocking receive after calling the function

// Closing a Channel
// Not necessary to close every channel
// Only close to indicate that no more messages will be sent to the channel
// Send operations will cause panic
// Receive operations will return zero value of the channel message type
// close(channel)

// Select
// Lets a goroutine wait on multiple processes, and will always block until one of the process can run
// Only used with channels and have similar syntax with switch
// Select is blocking, however if we use default it becomes a non-blocking routine and will automatically execute the default case
// select {case message := <- c}
