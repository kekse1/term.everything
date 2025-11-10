package pointerslices

import "slices"

func Contains[T comparable](slice []*T, item T) bool {
	for _, s := range slice {
		if s == nil {
			continue
		}
		if *s == item {
			return true
		}
	}
	return false
}

func DeleteFunc[T any](slice []*T, f func(T) bool) []*T {
	for i, s := range slice {
		if s == nil {
			continue
		}
		if f(*s) {
			return slices.Delete(slice, i, i+1)
		}
	}
	return slice
}

func Index[T comparable](slice []*T, item T) int {
	for i, s := range slice {
		if s == nil {
			continue
		}
		if *s == item {
			return i
		}
	}
	return -1
}

// Compares for value equality, not pointer equality
// But will return the index of nil if item is nil and there is a nil in the slice
func IndexOfItemOrNil[T comparable](slice []*T, item *T) int {
	for i, s := range slice {
		if s == nil && item == nil {
			return i
		}
		if s != nil && item != nil && *s == *item {
			return i
		}
	}
	return -1
}

func Delete[T any](slice []*T, start, end int) []*T {
	return slices.Delete(slice, start, end)
}

func Insert[T any](slice []*T, index int, values ...*T) []*T {
	return slices.Insert(slice, index, values...)
}
