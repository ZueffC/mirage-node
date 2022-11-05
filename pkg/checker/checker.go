package checker

func SizeCheck[T string](obj T, minSize, maxSize int) bool {
	if len(obj) >= minSize && len(obj) <= maxSize {
		return true
	}
	return false
}
