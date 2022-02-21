package utils

// SliceOrNil returns nil when the slice is empty, and the slice when it's not empty
func SliceOrNil(array []string) []string {
	if len(array) == 0 {
		return nil
	}
	return array
}
