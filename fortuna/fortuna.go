package main

import (
	"fmt"
	"time"
	"math/rand"
)
	
	var n int
    
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Ask Fortuna a question.")
	lens("Hello", "Goodbye", "yes", "no")
	roll()
}

func roll() {
	for i := 0; i < n; i++ {
		fmt.Println(lens(rand.Intn(4)))
	}
}
