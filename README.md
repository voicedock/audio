# audio
Audio utilities for golang

[![Go Report Card](https://goreportcard.com/badge/github.com/voicedock/audio)](https://goreportcard.com/report/github.com/voicedock/audio)
[![codecov](https://codecov.io/gh/voicedock/audio/branch/master/graph/badge.svg?token=53NH949TQY)](https://codecov.io/gh/voicedock/audio)
[![Go Reference](https://pkg.go.dev/badge/github.com/voicedock/audio.svg)](https://pkg.go.dev/github.com/voicedock/audio)
[![Release](https://img.shields.io/github/release/voicedock/audio.svg?style=flat-square)](https://github.com/voicedock/audio/releases/latest)

Written in Golang. Simple. Without 3rd party dependencies.

## Installation
```bash
go get github.com/voicedock/audio
```

## Usage
Reducing the sampling frequency from 48000 hertz to 16000 hertz (downsampling):
```go
inPcmData48000hz := []int16{/* ... */}
outPcmData16000hz := PcmDownsample[int16](inPcmData48000hz, 48000, 16000)
```
Converting PCM data represented by integers with bit depth 16 to float32:
```go
inPcmDataInt := []int16{/* ... */}
outPcmDataFloat := ConvertPcmIntToFloat[float32](16, inPcmDataInt)
```
Converting []int to []float32:
```go
ConvertNumbers[float32]([]int{1, 2, 3, 4, 5})
```

## Requirements
* Need at least `go1.19` or newer.

## Features
* Without external dependencies
* Simple
* Covered with tests

## Documentation
You can read package documentation [here](https://pkg.go.dev/github.com/voicedock/audio).

## Testing
Unit-tests:
```bash
go test -v -race ./...
```
Run linter:
```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.53 golangci-lint run -v
```

## CONTRIBUTE
* write code
* run `go fmt ./...`
* run all linters and tests (see above)
* create a PR describing the changes
