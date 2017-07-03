// +build linux darwin

package rawmap

import "os"
import "syscall"

// Returns the byte slice to the region of the memory-mapped file.
func MmapSharedReadonly(file *os.File, start, length int) ([]byte, error) {
	return syscall.Mmap(int(file.Fd()), int64(start), length, syscall.PROT_READ, syscall.MAP_SHARED)
}

// Unmaps the memory-mapped file pointed by the byte slice b.
func Munmap(b []byte) error {
	return syscall.Munmap(b)
}
