package su

import (
	"fmt"
	"sort"
	"strings"
)

type iterator[T any] struct {
	pos int
	s   *su[T]
}

type su[T any] struct {
	slice []T
}

type Iterator[T any] interface {
	// MoveNext method returns boolean if has next data.
	MoveNext() bool

	// Currnet method returns a element index and value.
	// If slice have no data then index is -1.
	Current() (int, T)

	// Reset a iterator's position.
	Reset()

	// Remove current a element of slice.
	Remove()
}

type Su[T any] interface {
	// Append method appends elements to the end of a slice.
	Append(values ...T)

	// Prepend method appends elements to the begin of a slice.
	Prepend(values ...T)

	// Remove method removes element at the index of a slice.
	Remove(index int) T

	// RemoveRange method removes elements from the start index to the end index of a slice.
	RemoveRange(start, end int)

	// RemoveAll method removes all elements.
	RemoveAll()

	// Len returns a slice size.
	Len() int

	// IsEmpty returns whether a slice is empty.
	IsEmpty() bool

	// Ptr method returns pointer of a element.
	Ptr(index int) *T

	// GetSlice method returns a slice.
	GetSlice() []T

	// Get method returns a element.
	Get(index int) T

	// Set method set a element at the index.
	Set(index int, value T)

	// Join method joins element to string.
	Join(seperator ...string) string

	// String method returns a slice's information as a string type.
	String() string

	// Reverse method reverses a slice.
	Reverse()

	// Sort method sorts a slice what ever you want.
	Sort(f func(i, j int) bool)

	// TrueForAll method returns true if all elements are true.
	TrueForAll(f func(value T) bool) bool

	// Filter method returns filtered elements.
	Filter(f func(value T) bool) Su[T]

	// IndexOf method returns a index what you found. If not found it returns -1.
	IndexOf(f func(value T) bool) int

	// Some method returns true if any are found.
	Some(f func(value T) bool) bool

	// Map method change all element.
	Map(f func(value T) T)

	// Iterator method creates slice iterator.
	Iterator() Iterator[T]
}

func New[T any](values ...[]T) Su[T] {
	var result []T
	if len(values) > 1 {
		for _, v := range values {
			result = append(result, v...)
		}
	} else if len(values) == 1 {
		result = values[0]
	}

	return &su[T]{
		slice: result,
	}
}

func (s *su[T]) RemoveAll() {
	s.slice = []T{}
}

func (s *su[T]) Append(values ...T) {
	s.slice = append(s.slice, values...)
}

func (s *su[T]) Prepend(values ...T) {
	s.slice = append(values, s.slice...)
}

func (s *su[T]) Remove(index int) T {
	ret := s.slice[index]
	s.slice = append(s.slice[:index], s.slice[index+1:]...)
	return ret
}

func (s *su[T]) RemoveRange(start, end int) {
	s.slice = append(s.slice[:start], s.slice[end+1:]...)
}

func (s *su[T]) Len() int {
	return len(s.slice)
}

func (s *su[T]) IsEmpty() bool {
	return s.Len() <= 0
}

func (s *su[T]) Ptr(index int) *T {
	return &s.slice[index]
}

func (s *su[T]) Get(index int) T {
	return s.slice[index]
}

func (s *su[T]) Set(index int, value T) {
	s.slice[index] = value
}

func (s *su[T]) GetSlice() []T {
	return s.slice
}

func (s *su[T]) Join(seperator ...string) string {
	var sep string
	if len(seperator) <= 0 {
		sep = ","
	} else {
		sep = seperator[0]
	}
	var ret strings.Builder
	size := s.Len()
	for i := 0; i < size; i++ {
		if i == size-1 {
			fmt.Fprintf(&ret, "%v", s.slice[i])
		} else {
			fmt.Fprintf(&ret, "%v%s", s.slice[i], sep)
		}
	}
	return ret.String()
}

func (s *su[T]) Reverse() {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - i - 1
		s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
	}
}

func (s su[T]) String() string {
	return fmt.Sprintf("[%s]", s.Join(", "))
}

func (s *su[T]) Sort(f func(i, j int) bool) {
	sort.Slice(s.slice, func(k, l int) bool {
		return f(k, l)
	})
}

func (s *su[T]) TrueForAll(f func(value T) bool) bool {
	for _, v := range s.slice {
		if !f(v) {
			return false
		}
	}
	return true
}

func (s *su[T]) Filter(f func(value T) bool) Su[T] {
	ret := New[T]()
	for _, element := range s.slice {
		if f(element) {
			ret.Append(element)
		}
	}
	return ret
}

func (s *su[T]) IndexOf(f func(value T) bool) int {
	for i, element := range s.slice {
		if f(element) {
			return i
		}
	}
	return -1
}

func (s *su[T]) Some(f func(value T) bool) bool {
	for _, element := range s.slice {
		if f(element) {
			return true
		}
	}
	return false
}

func (s *su[T]) Map(f func(value T) T) {
	for i := range s.slice {
		s.slice[i] = f(s.slice[i])
	}
}

func (s *su[T]) Iterator() Iterator[T] {
	return &iterator[T]{
		pos: -1,
		s:   s,
	}
}

func (i *iterator[T]) MoveNext() bool {
	i.pos += 1
	return i.s.Len() > i.pos
}

func (i *iterator[T]) Current() (int, T) {
	if i.pos >= i.s.Len() || i.pos < 0 {
		var ret T
		return -1, ret
	}
	return i.pos, i.s.slice[i.pos]
}

func (i *iterator[T]) Reset() {
	i.pos = -1
}

func (i *iterator[T]) Remove() {
	i.s.Remove(i.pos)
	i.pos--
}
