package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// ReadInputString reads a string from Stdin and puts that value in s
func ReadInputString(displayText string, s *[]byte) {
	fmt.Print(displayText)
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	*s = in.Bytes()
}

// ReadValidatedInputString reads lines from Stdin until a pattern is achieved, then stores in s
func ReadValidatedInputString(displayText string, s *string, pattern string, minLength int, maxLength int) error {
	var buff []byte

	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	for {
		ReadInputString(displayText, &buff)
		matched := re.Match(buff)

		switch {
		case !matched:
			fmt.Printf("Invalid value: %s\n", buff)
		case len(buff) < minLength:
			fmt.Printf("Input must be at least %d character(s) long\n", minLength)
		case len(buff) > maxLength:
			fmt.Printf("Input cannot be longer than %d character(s)\n", maxLength)
		default:
			*s = string(buff)
			return nil
		}
	}
}

// ReadBool asks user to input y/n and returns a bool
func ReadBool(displayText string) (bool, error) {
	var res string
	text := fmt.Sprintf("%s [y/n]: ", displayText)

	err := ReadValidatedInputString(text, &res, "[yn]", 1, 1)

	switch {
	case err != nil:
		return false, err
	case res == "y":
		return true, nil
	case res == "n":
		return false, nil
	default:
		return false, fmt.Errorf("Cannot convert input to bool: %s", res)
	}
}

// ReadInputStringHideInput reads input from user hiding Stdin chars
func ReadInputStringHideInput(displayText string) ([]byte, error) {
	fmt.Print(displayText)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return nil, err
	}

	return bytePassword, nil
}
