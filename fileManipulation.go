package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func addToFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFileSample(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)

	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var message string

	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF {
			break
		}
	}
	return message, nil
}

func main() {
	//createNewFile("sample.log", "this is sample.log")
	//sample, err := readFileSample("fileManipulation.go")
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(sample)
	//}
	addToFile("sample.log", "\nthis is new message")
}
