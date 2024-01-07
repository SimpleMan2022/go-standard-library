package main

import (
	"fmt"
	"os"
)

func main() {
	argumen := os.Args
	for _, args := range argumen {
		fmt.Println(args)
	}

	res, err := os.Hostname()
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
