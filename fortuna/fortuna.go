//Emulate Magic 8-ball
package main

import (
	"fmt"
	"time"
	"math/rand"
)
	var n = 20
	var q string
func main() {
Ask:
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(n)
	answer := []string{"It is certain", "It is decidedly so", "Without a doubt", "Yes, definitely", "You may rely on it", "As I see it, yes", "Most likely", "Outlook good", "Yes", "Signs point to yes", "Reply hazy try again", "Ask again later", "Better not tell you now", "Cannot predict now", "Concentrate and ask again", "Don't count on it", "My reply is no", "My sources say no", "Outlook not so good", "Very doubtful"}	
	fmt.Println("Ask Fortuna a question.")
	fmt.Scanln(&q)
	fmt.Println(answer[r])
	goto Ask
}
