package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		a, err := getInput("Enter value A: ")
		if err != nil {
			exitWithError(err)
		}

		b, err := getInput("Enter value B: ")
		if err != nil {
			exitWithError(err)
		}

		if a < 1 || b < 1 {
			break
		}

		fmt.Printf("GCD: %d\n", gcd(a, b))
		fmt.Printf("LCM: %d\n", lcm(a, b))
	}
}

func getInput(prompt string) (int, error) {
	var value int
	fmt.Print(prompt)
	_, err := fmt.Scanln(&value)
	return value, err
}

func exitWithError(err error) {
	fmt.Printf("Invalid input: %v. Exiting...\n", err)
	os.Exit(1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return b / gcd(a, b) * a
}
