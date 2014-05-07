//Decreases a population size by half over generations until 42
package main

import (
    "fmt"
    "math"
)
	var pop, g float64
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v > lim {
        return v
    }
    return lim
}



func main() {
	fmt.Println("Please enter a starting Population:")
	fmt.Scanln(&pop)
	fmt.Println("Please enter generations:")
	fmt.Scanln(&g)
	result := pop / pow(2, g, 2) 
	if result > 42 {
			fmt.Println(result)
	} else if result < 42 {
		fmt.Println("Reached 42")
		}
	}
