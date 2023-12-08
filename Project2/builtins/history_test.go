package builtins_test

import (
	"github.com/dangxnguyen/CSCE4600/Project2/builtins"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestSaveToHistoryAndShowHistory(t *testing.T) {
	// Create a temporary directory and change to it
	tmpDir, err := ioutil.TempDir("", "testhistory")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temporary directory: %v", err)
	}
	defer os.Chdir(originalDir) // Change back to the original directory

	// Test data
	commands := []string{"echo foo", "cd /tmp", "ls -l"}

	// Write commands to history
	for _, cmd := range commands {
		if err := builtins.SaveToHistory(cmd); err != nil {
			t.Errorf("SaveToHistory() failed for command '%s': %v", cmd, err)
		}
	}

	// Capture output of ShowHistory
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	if err := builtins.ShowHistory(); err != nil {
		t.Errorf("ShowHistory() failed: %v", err)
	}

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	output := string(out)
	// Check if all commands are in the output
	for i, cmd := range commands {
		expected := strings.TrimSpace(cmd) // Trimming spaces for safe comparison
		if !strings.Contains(output, expected) {
			t.Errorf("Expected command '%s' at line %d not found in history output", expected, i+1)
		}
	}
}


