package utils

func NewStringPointer(s string) *string {
	return &s
}

func NewStringSlice() []string {
	return make([]string, 0)
}
