package main

import (
	"fmt"
	"github.com/vadimturkov/go-algorithm-projects/util"
	"time"
)

func sieveOfEratosthenes(max int) []bool {
	sieve := make([]bool, max+1)

	sieve[2] = true
	for i := 3; i <= max; i += 2 {
		sieve[i] = true
	}

	for i := 2; i*i <= max; i++ {
		if sieve[i] == true {
			for j := i * i; j <= max; j += i {
				sieve[j] = false
			}
		}
	}

	return sieve
}

func printSieve(sieve []bool) {
	fmt.Printf("%d ", 2)
	for i := 3; i <= len(sieve)-1; i += 2 {
		if sieve[i] == true {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func sieveToPrimes(sieve []bool) []int {
	primes := []int{2}

	for i := 3; i < len(sieve)-1; i += 2 {
		if sieve[i] == true {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	max := util.GetIntInputOrExit("Max: ")

	start := time.Now()
	sieve := sieveOfEratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		printSieve(sieve)

		primes := sieveToPrimes(sieve)
		fmt.Println(primes)
	}
}
