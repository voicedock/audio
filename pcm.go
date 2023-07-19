package audio

import "math"

// DownsamplePcm downsamples the sample rate to the given value using averaging values.
// An example of reducing the sampling frequency from 48000 hertz to 16000 hertz:
// PcmDownsample[int16]([]int16{...}, 48000, 16000)
//
// Note: attempting to upsample will return the result unchanged
func DownsamplePcm[U, T Number](pcmData []T, srcSampleRate, dstSampleRate int) []U {
	sampleRateRatio := srcSampleRate / dstSampleRate
	if sampleRateRatio <= 1 {
		return ConvertNumbers[U](pcmData)
	}

	newPcmData := make([]U, len(pcmData)/sampleRateRatio)
	var offsetResult = 0
	var offsetBuffer = 0
	for offsetResult < len(newPcmData) {
		var nextOffsetBuffer = int(math.Round(float64(offsetResult+1) * float64(sampleRateRatio)))
		// Use average value of skipped samples
		var accum float64
		var count float64
		for i := offsetBuffer; i < nextOffsetBuffer && i < len(pcmData); i++ {
			accum += float64(pcmData[i])
			count++
		}

		newPcmData[offsetResult] = U(accum / count)
		offsetResult++
		offsetBuffer = nextOffsetBuffer
	}

	return newPcmData
}

// ConvertPcmIntToFloat converts PCM data represented by integers to floating point format.
// An example of converting PCM data represented by integers with bit depth 16 to float32:
// ConvertPcmIntToFloat[float32](16, []int16{...})
func ConvertPcmIntToFloat[U FloatNumber, T IntNumber](sourceBitDepth int, data []T) []U {
	out := make([]U, len(data))
	factor := math.Pow(2, float64(sourceBitDepth)-1)

	for i := 0; i < len(data); i++ {
		out[i] = U(float64(data[i]) / factor)
	}

	return out
}
