//go:build go1.18 && !go1.23
// +build go1.18,!go1.23

package tools

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 | uintptr |
	float32 | float64 | string
}