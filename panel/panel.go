//Prompt screen sample 
package main

import (
	"fmt"
)
	var name string
	var c, c2 string
func main() {
	fmt.Println("Welcome to the Panel!")
	fmt.Println("What's your name?")
	fmt.Scanln(&name)
	fmt.Println("Okay,", name, ",what would you like to do?")
	choices := []string {"look", "learn", "words", "market"}
	fmt.Scanln(&c)
	if c == choices[0] {
		fmt.Println("I see a red door and I want to paint it black.")
	} else if c == choices[1] {
		fmt.Println("Study Go or Python")
	} else if c == choices[2] {
		fmt.Println("Run wiki.py")
	} else if c == choices[3] {
		fmt.Println("Bling bling")
	} else { 
		fmt.Println("End of Line")
	}
	fmt.Scanln(&c2)
}
