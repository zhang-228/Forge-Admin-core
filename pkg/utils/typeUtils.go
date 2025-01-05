package utils

import "fmt"

// IsZeroValue checks if a value is the zero value of its type.
func IsZeroValue(v any) bool {
	switch v.(type) {
	case string:
		return v == ""
	case []string, []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float32, []float64:
		val := fmt.Sprintf("%v", v)
		return val == "[]" || val == "nil"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v == 0
	case float32, float64:
		return v == 0.0
	default:
		return true
	}
}
