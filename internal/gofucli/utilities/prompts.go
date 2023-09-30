package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Confirm(description string) bool {
	reader := bufio.NewReader(os.Stdin)
	tries := 3

	for ; tries > 0; tries-- {
		fmt.Printf("%s\nAre you sure you want to continue? [y/N]: ", description)

		answer, err := reader.ReadString('\n')
		if err != nil {
			return false
		}
		answer = strings.ToLower(strings.TrimSpace(answer))

		switch answer {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		}
	}

	return false
}
