package cli

import (
	"bufio"
	"fmt"
	"os"
)

func Ask(question string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s: ", question)

	if scanner.Scan() {
		input := scanner.Text()
		return input, nil
	}

	return "", scanner.Err()
}
