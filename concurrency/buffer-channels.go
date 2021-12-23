package concurrency

// There are two types of channels in Go: Buffered and Unbuffered
// channel := make(chan type) => Unbuffered Channel
// channel := make(chan type, capacity) => Buffered Channel

// Unbuffered Channel => Synchronous Channels
// Block the channel and the sending routine execution until the main routine receives the message from the channel
// Closed channel is unbuffered channel
// Literally waits for the main routine

// Buffered Channel => Asynchronous Channels
// Basically channel that uses buffer for its messages
// The sender routine will always send the message to the channel until the buffer is full
// If the buffer is full, then the main routine will block, execute and receive from the channel
// Apply FIFO messaging
