package concurrency

// WaitGroups
// WaitGroups can block and wait the program until all routines in that WaitGroups finished executing
// Always pass the waitGroup as a Pointer to the routine function. Also call waitGroup.Done() to indicate that the function is finished
// func routine(waitGroup *sync.WaitGroup) {defer waitGroup.Done()}
// In the spawner function => var waitGroup sync.WaitGroup
// waitGroup.Add(numberOfRoutinesToWait)
// go routine(&waitGroup)
// waitGroup.Wait() => wait for the execution of routines to finish
