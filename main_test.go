package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 10ms
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}

func Test_handleInput(t *testing.T) {
	// Capture the working directory before the test cases loop
	// wd, err := os.Getwd()
	// if err != nil {
	// 	t.Fatalf("Failed to get current working directory: %v", err)
	// }

	cases := []struct {
		name       string
		input      string
		wantOutput string
		wantErr    bool
		wantExit   bool
	}{
		{
			name:       "Echo Command",
			input:      "echo hello world",
			wantOutput: "hello world\n",
			wantErr:    false,
		},
		{
			name:     "Exit Command",
			input:    "exit",
			wantExit: true,
			wantErr:  false,
		},
		{
			name:       "Change Directory Command",
			input:      "cd ..",
			wantOutput: "",
			wantErr:    false,
		},
		{
			name:       "Environment Variables Command",
			input:      "env",
			wantOutput: "", 
			wantErr:    false,
		},
		{
			name:       "Help Command",
			input:      "help",
			wantOutput: "",
			wantErr:    true,
		},
		{
			name:       "History Command",
			input:      "history",
			wantOutput: "", 
			wantErr:    false,
		},
		{
			name:       "Print Working Directory Command",
			input:      "pwd",
			wantOutput: "", 
			wantErr:    false,
		},
		{
			name:       "Source Command",
			input:      "source somefile",
			wantOutput: "", 
			wantErr:    true, 
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
				var output bytes.Buffer
				exitChan := make(chan struct{}, 1)

				err := handleInput(&output, tc.input, exitChan)

				if tc.wantErr {
						assert.Error(t, err)
				} else {
						assert.NoError(t, err)
						assert.Contains(t, output.String(), tc.wantOutput, "Output for '%s' did not contain expected text", tc.name)
				}

				if tc.wantExit {
						_, ok := <-exitChan
						assert.True(t, ok, "Expected exit signal for '%s' was not received", tc.name)
				} else {
						assert.Empty(t, exitChan, "Exit channel for '%s' should be empty", tc.name)
				}
		})
	}
}


func Test_executeCommand(t *testing.T) {
	// Test a command that should succeed
	err := executeCommand("echo", "hello")
	require.NoError(t, err)

	// Test a command that should fail
	err = executeCommand("some-nonexistent-command")
	require.Error(t, err)
}