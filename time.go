package main

import (
	"fmt"
	"time"
)

func main() {
	jam := time.Now().Hour()
	menit := time.Now().Minute()
	detik := time.Now().Second()

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Printf("pukul %d:%d:%d\n", jam, menit, detik)
	fmt.Println(utc)
	fmt.Println(parse)

	durasi1 := time.Second * 100
	durasi2 := time.Millisecond * 10

	fmt.Println(durasi1)
	fmt.Println(durasi2)
}
