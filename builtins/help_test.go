package builtins_test

import (
	"bytes"
	"testing"

	"github.com/dangxnguyen/CSCE4600/builtins"
)

func TestHelp(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name: "No arguments",
			args: []string{},
			expected: "Available commands:\n" +
				"cd: Change the current directory\n" +
				"env: List environment variables\n" +
				"echo: Display a line of text\n" +
				"pwd: Print the current working directory\n" +
				"source: Execute commands from a file\n" +
				"help: Display help for available commands\n" +
				"history: Display the command history\n" +
				"exit: Exit the shell\n",
		},
		{
			name:     "With argument",
			args:     []string{"cd"},
			expected: "Description for 'cd': Change the current directory\n",
		},
		{
			name:     "Unknown command",
			args:     []string{"foo"},
			expected: "No help available for 'foo'\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			builtins.Help(&buf, test.args...)
			output := buf.String()

			if output != test.expected {
				t.Errorf("Expected: %q, Got: %q", test.expected, output)
			}
		})
	}
}
