package problems

import (
	"math"
)

func PrimesBelow(n int) []int {
	primes := []int{2}
	for i := 3; i < n; i += 2 {
		isPrime := true
		sqrti := int(math.Sqrt(float64(i)))
		for _, prime := range primes {
			if prime > sqrti {
				break
			}
			if i%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func GetAnswer3() int {
	primes := PrimesBelow(int(math.Sqrt(600851475143)))
	greatestFactor := 0
	for _, prime := range primes {
		if 600851475143%prime == 0 {
			greatestFactor = prime
		}
	}

	return greatestFactor
}
