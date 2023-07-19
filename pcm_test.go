package audio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertPcmIntToFloat(t *testing.T) {
	assert.Equal(
		t,
		[]float32{3.0517578e-05, 0.003753662, 0.03765869, 0.3767395},
		ConvertPcmIntToFloat[float32](16, []int16{1, 123, 1234, 12345}),
	)
}

func TestDownsamplePcm(t *testing.T) {
	// positive test
	assert.Equal(
		t,
		[]int16{1, 3, 5, 7},
		DownsamplePcm[int16]([]int16{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2),
	)

	// negative test
	assert.Equal(
		t,
		[]int16{1, 2, 3, 4, 5, 6, 7, 8},
		DownsamplePcm[int16]([]int16{1, 2, 3, 4, 5, 6, 7, 8}, 2, 4),
	)
}
