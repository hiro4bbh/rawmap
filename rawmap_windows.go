package rawmap

import "os"
import "syscall"
import "unsafe"

// Returns the byte slice to the region of the memory-mapped file.
func MmapSharedReadonly(file *os.File, start, length int) ([]byte, error) {
	fmap, err := syscall.CreateFileMapping(syscall.Handle(file.Fd()), nil, syscall.PAGE_READONLY, 0, 0, nil)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(fmap)
	ptr, err := syscall.MapViewOfFile(fmap, syscall.FILE_MAP_READ, uint32(start>>32), uint32(start), uintptr(length))
	if err != nil {
		return nil, err
	}
	// unsafe hack starts here
	bslice := Slice{ptr, length, length}
	return *(*[]byte)(unsafe.Pointer(&bslice)), nil
	// unsafe hack ends here
}

// Unmaps the memory-mapped file pointed by the byte slice b.
func Munmap(b []byte) error {
	// unsafe hack starts here
	return syscall.UnmapViewOfFile(uintptr(unsafe.Pointer(&b[0])))
	// unsafe hack ends here
}
