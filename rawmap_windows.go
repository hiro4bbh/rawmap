package rawmap

import "os"
import "syscall"
import "unsafe"

// Returns the byte slice to the region of the memory-mapped file with the specified protection mode.
func MmapShared(file *os.File, start, length int, protmode ProtectionMode) ([]byte, error) {
	prot_, access_ := 0, 0
	switch protmode {
	case PROTMODE_READONLY:
		prot_, access_ = syscall.PAGE_READONLY, FILE_MAP_READ
	case PROTMODE_READWRITE:
		prot_ = syscall.PAGE_READWRITE, FILE_MAP_READ | FILE_MAP_WRITE
	default:
		return nil, fmt.Errorf("unknown protmode: %#x", protmode)
	}
	fmap, err := syscall.CreateFileMapping(syscall.Handle(file.Fd()), nil, prot_, 0, 0, nil)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(fmap)
	ptr, err := syscall.MapViewOfFile(fmap, access_, uint32(start>>32), uint32(start), uintptr(length))
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
