package builtins_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/dangxnguyen/CSCE4600/builtins"
)

func TestPrintWorkingDirectory(t *testing.T) {
	// Capture stdout for testing output
	old := os.Stdout // Keep original stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // Set stdout to the writable pipe
	defer func() {
		os.Stdout = old // Reset stdout when done
	}()

	outCh := make(chan string)
	// Copy the output in a separate goroutine
	go func() {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		outCh <- buf.String()
	}()

	// Run the function you want to test
	_ = builtins.PrintWorkingDirectory()

	// Close the writable pipe to stop the goroutine
	_ = w.Close()

	output := <-outCh // Get captured output
	fmt.Println("Current Working Directory:", output)
	// Add assertions here to validate the output
}
