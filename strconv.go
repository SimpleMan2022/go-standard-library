package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, err := strconv.ParseBool("0")

	if err == nil {
		fmt.Println(b)
	} else {
		fmt.Println(err.Error())
	}

	i, err := strconv.Atoi("10")

	if err == nil {
		fmt.Println(i)
	} else {
		fmt.Println(err)
	}
}
