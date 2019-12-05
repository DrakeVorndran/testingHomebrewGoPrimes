package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var primes = []int{2, 3, 5}
	n := 7
	start := time.Now()

	for n < 1e9 {
		add := true
		for i := 0; float64(primes[i]) <= math.Sqrt(float64(n)); i++ {
			if n%primes[i] == 0 {
				add = false
				break
			}
		}
		if add {
			primes = append(primes, n)
			// fmt.Println(n)
		}
		n += 2
	}
	fmt.Println(primes)
	fmt.Println("num of primes:", len(primes))
	fmt.Println("time taken:", time.Since(start))
}
