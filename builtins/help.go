package builtins

import (
	"fmt"
	"io"
)

var commands = map[string]string{
	"cd":      "Change the current directory",
	"env":     "List environment variables",
	"echo":    "Display a line of text",
	"pwd":     "Print the current working directory",
	"source":  "Execute commands from a file",
	"help":    "Display help for available commands",
	"history": "Display the command history",
	"exit":    "Exit the shell",
}

func Help(w io.Writer, args ...string) {
	if len(args) > 0 {
		if desc, found := commands[args[0]]; found {
			fmt.Fprintf(w, "Description for '%s': %s\n", args[0], desc)
		} else {
			fmt.Fprintf(w, "No help available for '%s'\n", args[0])
		}
	} else {
		fmt.Fprintln(w, "Available commands:")
		for _, cmd := range []string{"cd", "env", "echo", "pwd", "source", "help", "history", "exit"} {
			if desc, found := commands[cmd]; found {
				fmt.Fprintf(w, "%s: %s\n", cmd, desc)
			}
		}
	}
}
