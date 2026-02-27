// Package slices основной пакет, который предоставляет доступ к разным реализациям массива.
//
//nolint:gosec
package slices

import "unsafe"

const (
	_startMem   = 0
	_defaultLen = 10
)

// LoadType определяет стратегию роста емкости.
type LoadType int

const (
	// AddOne - добавляем 1 элемент.
	AddOne LoadType = iota
	// AddHundred - добавляем 100 элементов.
	AddHundred
	// MultiplyByTwo - умножаем на 2.
	MultiplyByTwo
)

// Array базовый интерфейс для последующих имплементаций на основе фиксированного массива:
// - динамический массив
// - свободный массив
// - разреженный массив
// - параллельный массив
// - ассотиативный массив.
type Array[T any] interface {
	// Add добавляет элемент по указанному индексу
	Add(item T, index int) error

	// Get возвращает найденный элемент по указанному индексу
	Get(index int) (T, error)

	// Remove удаляет элемент по указанному индексу
	Remove(index int) error

	// Len возвращает кол-во элементов в массиве
	Len() int

	// IsEmpty проверяет не пустой ли массив
	IsEmpty() bool

	// Clear удаляет все элементы в массиве
	Clear()
}

// DynamicArray добавялет специфичные методы по работе с динамическим массивом.
type DynamicArray[T any] interface {
	Array[T]
	Capacity() int
	TrimToSize()
}

// SparseArray добавляет методы для работы с разреженным массивом.
type SparseArray[T any] interface {
	Array[T]
	NonZeroCount() int
	Compress()
}

// ParallelArray добавляет работу с мульти-парами.
type ParallelArray[T1, T2 any] interface {
	AddPair(first T1, second T2, index int) error
	GetFirst(index int) (T1, error)
	GetSecond(index int) (T2, error)
	Len() int
}

func loadTypeFactory(loadType LoadType) func(capacity int) int {
	switch loadType {
	case AddOne:
		return func(capacity int) int {
			return capacity + 1
		}
	case AddHundred:
		return func(capacity int) int {
			return capacity + 100
		}
	case MultiplyByTwo:
		return func(capacity int) int {
			if capacity == 0 {
				return 1
			}
			return capacity * 2
		}
	default:
		return func(capacity int) int {
			if capacity == 0 {
				return 1
			}
			return capacity * 2
		}
	}
}

func memcpy(dst, src unsafe.Pointer, size uintptr) {
	for i := range size {
		srcByte := *(*byte)(unsafe.Pointer(uintptr(src) + i))
		dstByte := (*byte)(unsafe.Pointer(uintptr(dst) + i))
		*dstByte = srcByte
	}
}

func memset(ptr unsafe.Pointer, value byte, size uintptr) {
	for i := range size {
		srcByte := (*byte)(unsafe.Pointer(uintptr(ptr) + i))
		*srcByte = value
	}
}
