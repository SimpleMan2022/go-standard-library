package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`a([a-z])i`)
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("ari"))
	fmt.Println(regex.MatchString("ado"))
	fmt.Println(regex.FindAllString("adi ari ani aka amar andi ati alo", 10))
}
