package utils

func IntOr(i *int, fallback int) int {
	if i == nil {
		return fallback
	}
	return *i
}
