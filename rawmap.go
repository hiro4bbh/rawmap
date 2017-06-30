// Package rawmap provides memory-mapped file functionalities.
package rawmap

// Slice is the data structure used in golang runtime.
type Slice struct {
	Addr     uintptr
	Len, Cap int
}
