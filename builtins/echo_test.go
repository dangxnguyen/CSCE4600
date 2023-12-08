package builtins_test

import (
	"bytes"
	"testing"

	"github.com/dangxnguyen/CSCE4600/builtins"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		expected     string
		expectedErr  error
	}{
		{
			name:     "Test single word",
			args:     []string{"Hello"},
			expected: "Hello\n",
		},
		{
			name:     "Test multiple words",
			args:     []string{"Hello", "World"},
			expected: "Hello World\n",
		},
		{
			name:     "Test empty arguments",
			args:     []string{},
			expected: "\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Redirect output to a buffer
			var output bytes.Buffer

			// Call the Echo function with test arguments and buffer as writer
			err := builtins.Echo(&output, test.args...)
			if err != nil {
				if test.expectedErr == nil {
					t.Fatalf("Echo() unexpected error: %v", err)
				}
			}

			// Check if the output matches the expected output
			if got := output.String(); got != test.expected {
				t.Errorf("Echo() = %q, expected %q", got, test.expected)
			}
		})
	}
}
