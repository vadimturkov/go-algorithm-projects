package main

import (
	"fmt"
	"github.com/vadimturkov/go-algorithm-projects/util"
)

func main() {
	for {
		a := util.GetIntInputOrExit("Enter value A: ")
		b := util.GetIntInputOrExit("Enter value B: ")

		if a < 1 || b < 1 {
			break
		}

		fmt.Printf("GCD: %d\n", gcd(a, b))
		fmt.Printf("LCM: %d\n", lcm(a, b))
	}
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
