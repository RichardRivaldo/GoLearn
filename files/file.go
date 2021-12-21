package files

// File interface in Go is platform-independent
// UNIX-like design and Go error handling
// Usually use OS Package, io, iotuil or others

// Create
// var newFile *os.File
// var err error
// newFile, err = os.Create("name.extension")

// Truncate
// Cut the file up to n bytes, clear the file if size = 0
// os.Truncate("file", size)

// Open
// os.Open("file") => Read Mode
// os.OpenFile("file", MODE | os.O_APPEND, PERMISSION)
// PERMISSION = 0644

// File Info
// var fileInfo os.FileInfo
// fileInfo, err = os.Stat("file")
// fileInfo.name(), fileInfo.Size(), fileInfo.ModTime(), fileInfo.Mode(), fileInfo.IsDir()

// Move or Rename
// os.Rename(oldPath, newPath)

// Remove
// os.Remove("file")

// Writing Bytes
// file.Write("file", bytesSlice)
// ioutil.WriteFile("file", bytesSlice, PERMISSIOn) => Handle create, open, write, and close

// Writing with Buffers - bufio
// writer = bufio.NewWriter(file)
// Available Buffer Size => writer.Available() => default 4MB = 4096 bytes
// written, err := writer.Write(bytesSlice) or writer.WriteString(string)
// Unflushed Buffer Size => writer.Buffered()
// Flush / apply the buffer content to the file => writer.Flush()
// Reset buffer content => writer.Reset(writer)

// Read based on size of the slice
// Read a file into slice of bytes
// io.readFull(file, targetSlice) => bytesRead, err

// Entire content
// ioutil.ReadAll(file) => str, err
// ioutil.ReadFile("file") => byteSlice, err => handles opening and closing

// Reading delimited file content
// scanner := bufio.NewScanner(file)
// scanner.Split(bufio.ScanRunes/Word) => delimiters if needed
// success := scanner.Scan()
// for scanner.Scan() {scanner.Text()}
// Most similar to Python file handling :)

// Close
// file.Close()
// defer file.Close()

// os > ioutil / other packages because of size comparisons
// bufio => provides buffer writer to minimize disk I/0 for many operations before dumping it to a file
