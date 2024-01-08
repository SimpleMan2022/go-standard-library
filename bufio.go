package main

import (
	"bufio"
	"os"
)

func main() {
	// reader
	//input := strings.NewReader("this is long string\nwith new line\n")
	//
	//reader := bufio.NewReader(input)
	//
	//for {
	//	line, _, err := reader.ReadLine()
	//
	//	if err == io.EOF {
	//		break
	//	} else {
	//		fmt.Println(string(line))
	//	}
	//}

	// writer
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("nama saya adit\n")
	writer.WriteString("saya semester 3\n")
	writer.WriteString("saya tinggal di binjai\n")
	writer.Flush()
}
