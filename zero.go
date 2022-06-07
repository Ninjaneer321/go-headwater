package headwater

// Get the zero value of type T
func GetZero[T any]() T {
	var zero T
	return zero
}
