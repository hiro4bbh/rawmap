// Package rawmap provides memory-mapped file functionalities.
package rawmap

import "os"

type ProtectionMode int

// The list of supported protection modes.
const (
	PROTMODE_READONLY ProtectionMode = iota
	PROTMODE_READWRITE
)

// Returns the corresponding flag used in os.OpenFile.
func (protmode ProtectionMode) OpenFlag() int {
	switch protmode {
	case PROTMODE_READONLY:
		return os.O_RDONLY
	case PROTMODE_READWRITE:
		return os.O_RDWR
	default:
		return 0
	}
}

// Slice is the data structure used in golang runtime.
type Slice struct {
	Addr     uintptr
	Len, Cap int
}
