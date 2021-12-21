package datstruct

// Map

// Key-Value, unordered data structure
// Operations and lookups are fast with constant time
// Unique keys
// Key = comparable types that can use ==, float is not recommended
// Can only compare map with nil just like slice
// Zero-Valued to nil if haven't yet been initialized
// Get will return zero-value of the type if the key doesn't exist

// Declaration
// Cannot populate the map if it is not initialized yet
// var m map[keyType]valType

// Operations
// len(map) => number of k-v pairs
// Comma OK Idiom to check when getting a value by a key from a map
// delete(map, key), no need to check if the key exists in the map
// fmt.Sprintf(map1) == fmt.Sprintf(map2) if both key and value type are strings

// Map Header
// A map refers to a map header stored in the memory
// = will just copy the reference of the map, and changing one will change others
// Can only manually copy each k-v of the reference map
