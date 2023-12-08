package builtins_test

import (
	"github.com/dangxnguyen/CSCE4600/Project2/builtins"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSource(t *testing.T) {
	// Helper function to create a temporary file with content
	createTempFileWithContent := func(content string) (string, error) {
		tmpFile, err := os.CreateTemp("", "testfile-")
		if err != nil {
			return "", err
		}
		defer tmpFile.Close()

		if _, err := tmpFile.WriteString(content); err != nil {
			return "", err
		}

		return tmpFile.Name(), nil
	}

	tests := []struct {
		name         string
		fileContent  string
		args         []string
		wantOutput   string
		expectError  bool
	}{
		{
			name:        "non-existent file",
			expectError: true,
		},
		{
			name:        "empty file",
			fileContent: "",
			wantOutput:  "",
		},
		{
			name:        "file with single command",
			fileContent: "echo hello",
			wantOutput:  "echo hello",
		},
		{
			name:        "file with multiple commands",
			fileContent: "echo hello\necho world",
			wantOutput:  "echo hello\necho world",
		},
		{
			name:        "file with command and extra args",
			fileContent: "echo",
			args:        []string{"hello", "world"},
			wantOutput:  "echo hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fileName string
			var err error

			// Handle non-existent file case
			if tt.name == "non-existent file" {
				fileName = "nonexistentfile"
			} else {
				fileName, err = createTempFileWithContent(tt.fileContent)
				if err != nil {
					t.Fatalf("Failed to create temp file: %v", err)
				}
				defer os.Remove(fileName)
			}

			// Capture the output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err = builtins.Source(fileName, tt.args...)
			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = oldStdout

			if (err != nil) != tt.expectError {
				t.Fatalf("Source() error = %v, expectError %v", err, tt.expectError)
			}

			if gotOutput := strings.TrimSpace(string(out)); gotOutput != tt.wantOutput {
				t.Errorf("Source() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
