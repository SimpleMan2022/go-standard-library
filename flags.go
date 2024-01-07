package main

import (
	"flag"
	"fmt"
)

func main() {
	var username = flag.String("username", "root", "put your username database")
	var password = flag.String("password", "root", "put your password database")
	var host = flag.String("host", "localhost", "put your host database")
	var port = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println(*username)
	fmt.Println(*password)
	fmt.Println(*host)
	fmt.Println(*port)

}
