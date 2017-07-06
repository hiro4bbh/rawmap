// +build linux darwin

package rawmap

import "fmt"
import "os"
import "syscall"

// Returns the byte slice to the region of the memory-mapped file with the specified protection mode.
func MmapShared(file *os.File, start, length int, protmode ProtectionMode) ([]byte, error) {
	prot_ := 0
	switch protmode {
	case PROTMODE_READONLY:
		prot_ = syscall.PROT_READ
	case PROTMODE_READWRITE:
		prot_ = syscall.PROT_READ | syscall.PROT_WRITE
	default:
		return nil, fmt.Errorf("unknown protmode: %#x", protmode)
	}
	return syscall.Mmap(int(file.Fd()), int64(start), length, prot_, syscall.MAP_SHARED)
}

// Unmaps the memory-mapped file pointed by the byte slice b.
func Munmap(b []byte) error {
	return syscall.Munmap(b)
}
