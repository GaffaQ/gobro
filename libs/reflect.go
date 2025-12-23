package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			fmt.Println(data)
			if result == false {
				return result
			}
		}
	}
	return result
}

type User struct {
	Name string `required:"true" max:"5"`
	Age  int    `required:"true" min:"17"`
}

func isValidUser(value any) bool {
	//valid := true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			v, ok := data.(string)
			if ok == true {
				if v != "" && len(v) <= 5 {
					return true
				} else {
					return false
				}
			} else {
				d, r := data.(int)
				if !r {
					return false
				}
				if d < 17 {
					return false
				} else {
					return true
				}
			}
		}
	}
	return true
}

func main() {
	//readField(Sample{"Eko"})
	//readField(Person{"Eko", "", ""})

	//person := Person{
	//	Name:    "ada",
	//	Address: "ada",
	//	Email:   "ada",
	//}
	//fmt.Println(IsValid(person))

	user := User{"Gaffa", 20}
	fmt.Println(isValidUser(user))
}
