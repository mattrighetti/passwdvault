package utils

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func ReadInputString(displayText string, s *string) {
	fmt.Print(displayText)
	fmt.Scanf("%s", s)
}

func ReadBool(displayText string) bool {
	var res string
	fmt.Printf("%s [y/n]: ", displayText)
	fmt.Scanf("%s", &res)
	if res != "y" {
		return false
	}

	return true
}

func ReadInputStringHideInput(displayText string) ([]byte, error) {
	fmt.Print(displayText)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}

	return bytePassword, nil
}
