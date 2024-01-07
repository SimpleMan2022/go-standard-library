package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max: "10"`
}

type Person struct {
	Name     string `required:"true" max: "10"`
	Email    string `required:"true" max: "10"`
	Password string `required:"true" max: "10"`
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result

}

func readFile(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println(valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "and type is", valueField.Type, "and tag:", valueField.Tag.Get("required"))
	}
}

func main() {
	//readFile(Sample{Name: "adit"})
	//readFile(Person{Name: "bagas", Age: 19})
	tes := isValid(Person{"adit", "saas", "aass"})
	fmt.Println(tes)
}
