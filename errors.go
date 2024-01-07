package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func GetUsersById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "adit" {
		return NotFoundError
	}
	return nil
}

func main() {
	result := GetUsersById("")
	if result != nil {
		if errors.Is(result, ValidationError) {
			fmt.Println(ValidationError)
		} else if errors.Is(result, NotFoundError) {
			fmt.Println(NotFoundError)
		} else {
			fmt.Println("unknown error")
		}
	} else {
		fmt.Println("sukses")
	}
}
