package audio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertNumbers(t *testing.T) {
	// []float32 to []int
	assert.Equal(t, []int{1, 2, 3, 4, 5}, ConvertNumbers[int]([]float32{1, 2, 3, 4, 5}))
	// []int to []float32
	assert.Equal(t, []float32{1, 2, 3, 4, 5}, ConvertNumbers[float32]([]int{1, 2, 3, 4, 5}))
}
