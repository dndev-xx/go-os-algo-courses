//nolint:gosec,revive
package slices

import (
	"errors"
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

type DArray[T any] struct {
	data     unsafe.Pointer
	length   int
	capacity int

	elemSize   uintptr
	growFactor func(int) int
}

func New[T any](l int, capacity int, loadFactory LoadType) *DArray[T] {
	if capacity <= 0 {
		capacity = _defaultLen
	}
	if l < 0 {
		l = 0
	}
	if l > capacity {
		l = capacity
	}

	var elem T
	elemSize := unsafe.Sizeof(elem)
	totalSize := elemSize * uintptr(capacity)
	data, err := syscall.Mmap(
		-1,
		_startMem,
		int(totalSize),
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to allocate memory: %v", err))
	}
	return &DArray[T]{
		data:       unsafe.Pointer(&data[_startMem]),
		length:     l,
		capacity:   capacity,
		elemSize:   elemSize,
		growFactor: loadTypeFactory(loadFactory),
	}
}

func (da *DArray[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= da.length {
		return zero, fmt.Errorf("index out of bounds: %d", index)
	}

	elemPtr := unsafe.Pointer(uintptr(da.data) + da.elemSize*uintptr(index))
	return *(*T)(elemPtr), nil
}

func (da *DArray[T]) Add(item T, index int) error {
	if !da.isOutOfRangeIndex(index) {
		return errors.New("index out of range")
	}

	if da.length >= da.capacity {
		da.grow()
	}
	for i := da.length; i > index; i-- {
		srcPtr := unsafe.Pointer(uintptr(da.data) + da.elemSize*uintptr(i-1))
		dstPtr := unsafe.Pointer(uintptr(da.data) + da.elemSize*uintptr(i))
		memcpy(dstPtr, srcPtr, da.elemSize)
	}

	elemPtr := unsafe.Pointer(uintptr(da.data) + da.elemSize*uintptr(index))
	*(*T)(elemPtr) = item
	da.length++

	return nil
}

func (da *DArray[T]) Remove(index int) error {
	if !da.isOutOfRangeIndex(index) {
		return errors.New("index out of range")
	}
	for i := range index {
		srcPtr := unsafe.Pointer(uintptr(da.data) + da.elemSize*uintptr(i+1))
		dstPtr := unsafe.Pointer(uintptr(da.data) + da.elemSize*uintptr(i))
		memcpy(dstPtr, srcPtr, da.elemSize)
	}

	if da.length > 0 {
		lstPtr := unsafe.Pointer(uintptr(da.data) + da.elemSize*uintptr(da.length-1))
		memset(lstPtr, 0, da.elemSize)
	}
	da.length--
	return nil
}

func (da *DArray[T]) String() string {
	slice := unsafe.Slice((*T)(da.data), da.length)
	return fmt.Sprintf("%v", slice)
}

func (da *DArray[T]) Len() int {
	return da.length
}

func (da *DArray[T]) Cap() int {
	return da.capacity
}

func (da *DArray[T]) Close() error {
	if da.data == nil {
		return nil
	}

	totalSize := da.elemSize * uintptr(da.capacity)
	data := unsafe.Slice((*byte)(da.data), int(totalSize))
	err := syscall.Munmap(data)
	if err == nil {
		da.data = nil
		da.length = 0
		da.capacity = 0
	}

	return err
}

func (da *DArray[T]) isOutOfRangeIndex(index int) bool {
	if index < 0 || index > da.length {
		return false
	}
	return true
}

func (da *DArray[T]) grow() {
	newCap := da.growFactor(da.capacity)

	if newCap <= da.capacity {
		newCap = da.capacity + 1
	}

	oldSize := da.elemSize * uintptr(da.capacity)
	newSize := da.elemSize * uintptr(newCap)

	data, err := syscall.Mmap(
		-1,
		_startMem,
		int(newSize),
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to allocate memory: %v", err))
	}
	newPtr := unsafe.Pointer(&data[0])
	if da.length > 0 && da.data != nil {
		memcpy(newPtr, da.data, oldSize)
	}

	if da.data != nil {
		oldData := unsafe.Slice((*byte)(da.data), int(oldSize))
		if err = syscall.Munmap(oldData); err != nil {
			log.Fatalf("failed to munmap: %v", err)
		}
	}

	da.data = newPtr
	da.capacity = newCap
}
