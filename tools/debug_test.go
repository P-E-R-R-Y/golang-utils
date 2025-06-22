package tools_test

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/P-E-R-R-Y/golang-utils/tools"
)

func TestTrace(t *testing.T) {
	// Backup original os.Stdout
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		outC <- buf.String()
	}()

	// Run the function that prints
	func() {
		defer tools.Trace("TestFunc")()
		time.Sleep(10 * time.Millisecond)
	}()

	// Restore original os.Stdout and close the writer
	_ = w.Close()
	os.Stdout = old

	// Read captured output
	output := <-outC

	if !strings.Contains(output, "TestFunc took") {
		t.Errorf("Trace output missing expected prefix, got: %q", output)
	}
}