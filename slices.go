package main

import (
	"fmt"
	"slices"
)

func main() {
	orang := []string{"adit", "nugroho", "rosyad"}
	nilai := []int{90, 98, 78, 88}

	fmt.Println(slices.Contains(orang, "nugroho"))
	fmt.Println(slices.Index(orang, "rosyad"))
	fmt.Println(slices.Min(nilai))
}
