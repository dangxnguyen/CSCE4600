package builtins

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Source(fileName string, args ...string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := scanner.Text()
		fullCommand := command
		if len(args) > 0 {
			// Append each argument individually
			fullCommand = fmt.Sprintf("%s %s", command, strings.Join(args, " "))
		}
		fmt.Println(fullCommand)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
