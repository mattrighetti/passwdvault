package utils

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// ReadInputString reads a string from Stdin and puts that value in s
func ReadInputString(displayText string, s *string) {
	fmt.Print(displayText)
	fmt.Scanf("%s", s)
}

// ReadBool asks user to input y/n and returns a bool
func ReadBool(displayText string) bool {
	var res string
	fmt.Printf("%s [y/n]: ", displayText)
	fmt.Scanf("%s", &res)
	if res != "y" {
		return false
	}

	return true
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
