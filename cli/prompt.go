package cli

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

var (
	stdinFd = int(os.Stdin.Fd())
)

func promptText(linePrefix string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(linePrefix)
	scanner.Scan()
	return scanner.Text()
}

func promptPassword(linePrefix string) (string, error) {

	fmt.Print(linePrefix)
	password, err := term.ReadPassword(stdinFd)

	if err != nil {
		return "", err
	}
	fmt.Print("\n")
	return string(password), nil
}
