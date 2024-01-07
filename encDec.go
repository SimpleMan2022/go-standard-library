package main

import (
	"encoding/csv"
	"os"
)

func main() {
	//encoded := base64.StdEncoding.EncodeToString([]byte("adit nugroho"))
	//fmt.Println(encoded)
	//
	//decoded, err := base64.StdEncoding.DecodeString(encoded)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(string(decoded))
	//}

	//csv reader

	//csvFile := "rosyad,aditya,nugroho\n" +
	//	"alex,bayu,andre\n" +
	//	"didit,maldi,suci"
	//reader := csv.NewReader(strings.NewReader(csvFile))
	//
	//for {
	//	record, err := reader.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(record)
	//}

	// csv writer
	csvWriter := csv.NewWriter(os.Stdout)
	_ = csvWriter.Write([]string{"adit", "nugroho", "mantap"})
	_ = csvWriter.Write([]string{"bayu", "anggara", "mantap"})
	_ = csvWriter.Write([]string{"tribudi", "kalia", "mantap"})

	csvWriter.Flush()
}
