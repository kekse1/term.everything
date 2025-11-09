package wayland

/*
#include "mmap.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type MemMapInfo struct {
	Bytes          []byte
	Addr           unsafe.Pointer
	Size           C.size_t
	FileDescriptor C.int
	UnMapped       bool
}

func NewMemMapInfo(fd int, size uint64) (MemMapInfo, error) {
	fdNum := C.int(fd)
	c_size := C.size_t(size)
	addr := C.mmap_fd(fdNum, c_size)
	if addr == C.map_failed() {
		return MemMapInfo{
			Addr:           addr,
			Size:           c_size,
			FileDescriptor: fdNum,
			UnMapped:       true,
		}, fmt.Errorf("failed to mmap fd %d", fdNum)
	}

	info := MemMapInfo{
		Addr:           addr,
		Size:           c_size,
		FileDescriptor: fdNum,
		UnMapped:       false,
	}
	info.UpdateBytes()
	return info, nil
}

// This happens automatically on creation and remap
func (m *MemMapInfo) UpdateBytes() {
	m.Bytes = unsafe.Slice((*byte)(m.Addr), m.Size)
}

func (m *MemMapInfo) Remap(newSize uint64) error {
	if m.UnMapped {
		return fmt.Errorf("cannot remap unmapped memory")
	}
	c_newSize := C.size_t(newSize)

	newAddr := C.remap(m.FileDescriptor, m.Addr, m.Size, c_newSize)
	if newAddr == C.map_failed() {
		m.UnMapped = true
		return fmt.Errorf("failed to remap memory")
	}

	m.Addr = newAddr
	m.Size = c_newSize
	m.UpdateBytes()
	return nil
}

func (m *MemMapInfo) Unmap() {
	if m.UnMapped {
		return
	}

	C.unmap(m.Addr, m.Size)
	m.UnMapped = true
	m.Bytes = nil
}
