package test

// ComparePointers values
func ComparePointers[T comparable](a, b *T) bool {
	if a == nil && b != nil ||
		a != nil && b == nil {

		return false
	}

	if a == nil && b == nil {
		return true
	}

	return *a == *b
}
