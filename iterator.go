package headwater

func ForEach[T any](array []T, callback func(item T)) {
	for _, item := range array {
		callback(item)
	}
}

func Map[T any, U any](array []T, callback func(item T) U) []U {
	var output []U = make([]U, len(array))
	for index, item := range array {
		result := callback(item)
		output[index] = result
	}
	return output
}

func Reduce[T any, U any](array []T, callback func(target U, item T) U, target U) U {
	var output U = target
	for _, item := range array {
		output = callback(output, item)
	}
	return output
}

func Filter[T any](array []T, callback func(item T) bool) []T {
	var output []T = make([]T, 0)
	for _, item := range array {
		valid := callback(item)
		if valid {
			output = append(output, item)
		}
	}
	return output
}

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
