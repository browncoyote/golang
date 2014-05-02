package main

import (
	"fmt"
	"net"
)
	 
	var name string
	var answer string
	var host string
func main() {
	fmt.Println("Welcome to a Go program called Name2.")
	fmt.Println("What's your name?")
	fmt.Scanln(&name)
	fmt.Println("Interesting name,", name, ",do you know what it means?")
Ask:
	fmt.Scanln(&answer)
	posAnswers := []string{"y", "Y", "yes", "Yes", "YES", "ja"}
	negAnswers := []string{"n", "N", "no", "No", "NO"}
	bsAnswers := []string{"blerg", "kinda", "pissoff", "smartass"}
	if containsString(posAnswers, answer) {
		fmt.Println("Oh really?")
		goto Ask
	} else if containsString(negAnswers, answer) {
		fmt.Println("Well let me look it up!")
		goto Query
	} else if containsString(bsAnswers, answer) {
		fmt.Println("Piss off type y/n, smart ass")
		goto Ask
	} else { 
		fmt.Println("Please type yes or no, then press Enter.")
		goto Ask
	}
Query:
	fmt.Println("Enter host:")
	fmt.Scanln(&host)
	fmt.Println("Scanning the internets.")
	ip, _ := net.LookupIP(host)
    fmt.Println(ip)
    goto Query
}
	
//two functions containsString and posString for study
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}
