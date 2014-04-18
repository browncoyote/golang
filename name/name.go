package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("What's your name?")
	reader := bufio.NewReader(os.Stdin)
	name,err := reader.ReadString('\n')
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return
	}
	
	fmt.Printf("Hi, %s!", name)
}
