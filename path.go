package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("belajar-go/tes.go"))
	fmt.Println(path.Base("belajar-go/tes.go"))
	fmt.Println(path.Ext("belajar-go/tes.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))

	fmt.Println(filepath.Dir("belajar-go/tes.go"))
	fmt.Println(filepath.Base("belajar-go/tes.go"))
	fmt.Println(filepath.Dir("belajar-go/tes.go"))
	fmt.Println(filepath.IsLocal("belajar-go/tes.go"))
	fmt.Println(filepath.IsAbs("belajar-go/tes.go"))

}
