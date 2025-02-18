package safetyper

func PtrToString(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}

func IsNilPtr[T any](v *T) bool {
	return v == nil
}
