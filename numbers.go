package audio

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type IntNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type FloatNumber interface {
	~float32 | ~float64
}

// ConvertNumbers casts a slice of numbers to a slice of numbers of another type.
// Example convert []int to []float32: ConvertNumbers[float32]([]int{1, 2, 3, 4, 5})
func ConvertNumbers[U, T Number](s []T) (out []U) {
	out = make([]U, len(s))
	for i := range s {
		out[i] = U(s[i])
	}
	return out
}
