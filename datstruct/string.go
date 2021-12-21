package datstruct

// String
// Double quoted or backticks (raw string), UTF-8 encoded by default
// Raw string won't have any backslash effect like newlines
// String is slice of bytes in Go, indexing will return ASCII integer
// String is immutable
// Every byte slices can be encoded in a string value
// Indexing a string will return byte not rune

// Rune (uint8) and Byte (int32)
// Byte and Rune => distinguish char and integer
// Doesn't have char, use both byte and rune to represent char values
// Byte => ASCII, Rune => UTF-8 Unicode characters
// Unicode characters size from 1-4 bytes

// Chars or Rune Literals
// 'x' => Unicode Code Points => numeric value of a rune literal

// Decoding Rune
// utf8.DecodeRuneInString => return the rune and the size of the decoded rune
// %c => Rune / Chars

// Rune Length
// utf8.RuneCountInString => Only returns the number of Rune
// len(string_of_runes) => returns the number of bytes used for the Rune

// Slicing Strings
// Can slice a string normally if it only contains ASCII
// Else, convert the string to rune first and then convert back to string

// Functions
// strings package
// strings.Contains(str, substr) => Check substring, returns true or false
// strings.ContainsAny(str, substr)
// strings.ContainsRune(str, rune)  => Check rune
// strings.Count(str, substr)  => Count the number of a substring inside a string
// If the substring in Count is "", then it will return 1 + number of runes
// strings.ToLower and strings.ToUpper: (str)
// strings.Repeat(str, n) => repeat a string n times

// strings.Replace(str, replaced, new, counts) => replace all replaced symbols with the new one, -1 to replace all
// strings.ReplaceAll(str, replaced, new)
// strings.Split(string, sep) => Split the string based on the separator. "" will split the string by each rune literal,
// will still return slice of strings
// strings.Join(sliceOfStrings, sep)
// strings.Fields(str) => return slice of splitted strings based on newline or spaces
// strings.TrimSpace(str) => remove all preceding and trailing spaces and newlines
// strings.Trim(str, symbols) => remove all preceding and trailing symbols

// Operators
// Equality => ==
// strings.EqualFold => Compare two strings, not case-sensitive
