package main

import (
	"flag"
	"fmt"
)

func main() {

	// package flag
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")  
	var host *string = flag.String("host", "localhost", "database host")  
	var port *int = flag.Int("port", 0, "database port")  

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

	// cara run => go run flag.go -username=rizky -password=bismillah -host=123.234.456.7 -port=233
}