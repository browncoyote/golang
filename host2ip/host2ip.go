package main

import (
	"fmt"
	"net"
)
	var host string
func main() {
Query:
	fmt.Println("Enter host:")
	fmt.Scanln(&host)
	fmt.Println("Scanning the internets.")
	ip, _ := net.LookupIP(host)
    fmt.Println(ip)
    goto Query
}
	
