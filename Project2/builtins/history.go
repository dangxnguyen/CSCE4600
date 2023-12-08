// builtins/history.go

package builtins

import (
	"bufio"
	"fmt"
	"os"
)

const historyFile = ".history"

func SaveToHistory(command string) error {
	file, err := os.OpenFile(historyFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, "%s\n", command); err != nil {
		return err
	}

	return nil
}

func ShowHistory() error {
	file, err := os.OpenFile(historyFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Command history:")
	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", count, scanner.Text())
		count++
	}

	return nil
}
