package main

import (
	"fmt"
	"github.com/vadimturkov/go-algorithm-projects/util"
)

func main() {
	for {
		num := util.GetIntInputOrExit("Enter value num: ")
		pow := util.GetIntInputOrExit("Enter value pow: ")
		mod := util.GetIntInputOrExit("Enter value mod: ")

		if num < 1 || pow < 1 || mod < 1 {
			break
		}

		fmt.Printf("(num ^ pow): %d\n", fastExp(num, pow))
		fmt.Printf("(num ^ pow %% mod): %d\n", fastExpMod(num, pow, mod))
	}
}

func fastExp(num, pow int) int {
	result := 1

	for pow > 0 {
		if pow%2 == 1 {
			result *= num
		}
		num *= num
		pow /= 2
	}

	return result
}

func fastExpMod(num, pow, mod int) int {
	result := 1

	for pow > 0 {
		if pow%2 == 1 {
			result = (result * num) % mod
		}
		num = (num * num) % mod
		pow /= 2
	}

	return result
}
