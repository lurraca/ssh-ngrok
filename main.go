package main

import (
	"fmt"
	"strings"

	flag "github.com/ogier/pflag"
)

//flags
var (
	username string
)

func main() {

	flag.Parse()
	fmt.Printf("$$$$$$$$$$$$$$$ %s", username)
	fmt.Println(username)

	if strings.TrimSpace(username) == "" {
		fmt.Println("Missing Username. Correct usage: ssh-ngrok -u <username-on-host>")
		return
	}
	fmt.Printf("ssh%s@ngrok.io -p 9000", username)
}

func init() {
	flag.StringVarP(&username, "username", "u", "", "username on the host machine")
}
