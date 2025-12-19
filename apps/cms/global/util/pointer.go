package util

// Ptr returns a pointer to the given value
func Ptr[T any](v T) *T {
	return &v
}

// PtrOrNil returns a pointer to the value if non-empty, otherwise nil
func PtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
