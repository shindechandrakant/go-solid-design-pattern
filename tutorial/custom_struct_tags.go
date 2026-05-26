package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name  string `validate:"min=2,max=3"`
	Email string `validate:"required,email"`
}

func validate(val interface{}) error {

	//v := reflect.ValueOf(val)
	t := reflect.ValueOf(val)
	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		fmt.Printf("%+v\n", field)
		tag := t.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		fmt.Printf("%+v\n", tag)
		rules := strings.Split(tag, ",")

		for _, rule := range rules {
			fieldName := t.Type().Field(i).Name
			switch {
			case strings.HasPrefix(rule, "min="):
				min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				if len(field.String()) < min {
					return fmt.Errorf("%s must be at least %d long", fieldName, min)
				}
			case strings.HasPrefix(rule, "max="):
				max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
				if len(field.String()) > max {
					return fmt.Errorf("%s must be at less than %d ", fieldName, max)
				}
			}

		}

	}
	fmt.Printf("%+v", t)
	return nil
}

//func main() {
//
//	user := User{
//		Name:  "Chandraknt",
//		Email: "Chandra@gmail.com",
//	}
//
//	//cmp.Compare()
//	//t := reflect.TypeOf(user)
//	fmt.Println(validate(user))
//	//
//	//for i := 0; i < t.NumField(); i++ {
//	//	field := t.Field(i)
//	//	tag := field.Tag.Get("validate")
//	//	fmt.Printf("%d. %v. (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
//	//}
//
//}
