package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "Adit Nugroho"
	kataLower := strings.ToLower(kata)
	kataTrim := strings.Trim(kata, "A")
	kataSplit := strings.Split(kata, "")

	fmt.Println(kataLower)
	fmt.Println(kataTrim)
	fmt.Println(kataSplit)
	fmt.Println(strings.Join(kataSplit, ""))
}
