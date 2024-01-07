package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (user User) walking() {
	fmt.Printf("%s sedang berjalan.", user.Name)
}

func main() {
	users1 := []User{
		{"alex", 19},
		{"Adit", 17},
		{"bayu", 18},
		{"dika", 20},
	}
	sort.Sort(UserSlice(users1))
	for _, users := range users1 {
		fmt.Println(users)
	}
}
