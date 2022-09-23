package pointer

// SPtr returns a pointer for the given string
func SPtr(s string) *string { return &s }

// F64Ptr returns a pointer fot the given float 64
func F64Ptr(f float64) *float64 { return &f }
