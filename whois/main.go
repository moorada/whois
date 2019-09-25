package main

import (
	"fmt"
	"os"

	"github.com/moorada/whois"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: \n go run main.go example.com [server]")
		return
	}

	var server string
	if len(os.Args) > 2 {
		server = os.Args[2]
	}

	result, err := whois.Whois(os.Args[1], server)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("result: ", result)
	}

	os.Exit(0)
}
