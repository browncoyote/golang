//Log growth of a population size

package main

import (
    "fmt"
    "time"
	"math/rand"
)
	var pop, g, sum int
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to Random Growth")
	fmt.Println("Please enter a population size:")
	fmt.Scan(&pop)
	fmt.Println("Enter generations:")
	fmt.Scan(&g)
	grow(g, pop)
}

func grow(g, pop int) {
	for i := 0; i < g; i++ {
		sum = pop - (((rand.Intn(pop))) - ((rand.Intn(pop))))
		sum += sum
		if sum > 42 {
			fmt.Println(sum)
	} else if sum < 1 {
			fmt.Println("Dead")
	}
}
}
