package util

import (
	"fmt"
	"os"
)

func GetIntInputOrExit(prompt string) int {
	var value int
	fmt.Print(prompt)

	_, err := fmt.Scanln(&value)
	if err != nil {
		fmt.Printf("Invalid input: %v. Exiting...\n", err)
		os.Exit(1)
	}

	return value
}
