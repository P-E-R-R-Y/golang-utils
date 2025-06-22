package tools

import (
	"time"
	"fmt"
)

// Trace logs the duration of a function.
func Trace(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}