package datstruct

// Slice
// Has dynamic length, must be same type
// Zero-valued to nil, unlike array to 0
// Can use key index for slice elements
// Iterate just like array with range or manual

// Declaration
// var slice []type => not initialized == nil
// var slice = []type{} => initialized != nil
// slice := []type{val1, val2, ...}
// make([]type, length, capacity)

// Operators
// Cannot compare slices with ==, can only compare to nil
// Can only check elements and length of the slices manually
// = => used to assign same-type slice to another slice

// Operations
// Append => append(slice, element1, element2, ...) => return new slice
// Extend => append(slice, slice...)
// Copy => copy(target, source) => return number of elements copied

// Backing Array and Slice Header
// Go saves the elements of a slice into a Backing Array, nil = no Backing Array
// Go implements the slice into Slice Header, slice representation on runtime
// Slice Header has length, capacity, and pointer to first element of the Backing Array
// Length is the number of elements used by the slice
// Capacity is the number of elements in the Backing Array
// IMPORTANT: Elements of an array are contiguously stored in the memory

// Slicing
// Both array and slice can be sliced, returning new slice
// Slicing can get all the elements in the backing array even if the slice index is out of range of its length
// IMPORTANT: Slicing refers to the same Backing Array, changing the slice means changing the Backing Array,
// thus changing the whole slices that refers the backing array

// Copying a Slice
// Use append with ellipsis, because append returns new slice
// append(target, source[slices]...)
// Can also use copy
// The elements in the copy destination is based on the length of the target slice

// Append
// Append is an expensive operation
// If a slice's capacity cannot hold the append result, Go will create new Backing Array and add more capacity
// The capacity is added to the power of 2 and will decrease at a later time

// Slice vs Array
// Slicing is usually cheaper than array
// Slices have smaller size than array
// unsafe.sizeof(object)
// Pointer + 2 * Integer = 4 + 16 = 20 bytes
