package utils

func ValueOr[T any](t *T, fallback T) T {
	if t == nil {
		return fallback
	}
	return *t
}

func Ptr[T any](t T) *T {
	return &t
}

func PtrSlice[T any](input []T) (output []*T) {
	for _, t := range input {
		output = append(output, &t)
	}
	return
}
