package slices

import "github.com/PlayerR9/GoSD/pkg"

func FilterNilValues[T pkg.Type](slice *Slice[T]) *Slice[T] {
	if slice == nil {
		return nil
	}

	return slice
}
