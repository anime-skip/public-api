package math

// BoundedInt returns the value bounded by the min/max value
func BoundedInt(val, min, max int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}
