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
func ReadInputString(displayText string, s *string) {
	fmt.Print(displayText)
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	*s = in.Text()
}

// ReadValidatedInputString reads lines from Stdin until a pattern is achieved, then stores in s
func ReadValidatedInputString(displayText string, s *string, pattern string) error {
	var temp string

	for {
		ReadInputString(displayText, &temp)
		matched, err := regexp.Match(pattern, []byte(temp))

		if err != nil {
			return err
		} else if !matched {
			fmt.Println("Invalid value: ", temp)
		} else {
			*s = temp
			return nil
		}
	}
}

// ReadBool asks user to input y/n and returns a bool
func ReadBool(displayText string) (bool, error) {
	var res string
	fmt.Printf("%s [y/n]: ", displayText)
	fmt.Scanf("%s", &res)
	if res != "y" {
		return false, nil
	}

	return true, nil
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
