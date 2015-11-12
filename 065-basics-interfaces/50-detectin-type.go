package main

import (
	"fmt"
	"reflect"
)

//string formatting - short and low footprint (not necessary to import reflect package)
func typeof1(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

//reflect package - when need more details about the type we have access to the full reflection capabilities
func typeof2(v interface{}) string {
	return reflect.TypeOf(v).String()
}

//type assertions - allows grouping types, for example recognize all int32, int64, uint32, uint64 types as "int"
func typeof3(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	//... etc
	default:
		_ = t
		return "unknown"
	}
}

func main() {
	b := true
	s := ""
	n := 1
	f := 1.0
	a := []string{"foo", "bar", "baz"}

	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())
	fmt.Println(reflect.ValueOf(n).Kind())
	fmt.Println(reflect.ValueOf(f).Kind())
	fmt.Println(reflect.ValueOf(a).Index(0).Kind())

	fmt.Printf("%+v ", typeof1(s))
	fmt.Printf("%+v ", typeof2(s))
	fmt.Printf("%+v ", typeof3(s))
	fmt.Printf("\n")

	fmt.Printf("%+v ", typeof1(n))
	fmt.Printf("%+v ", typeof2(n))
	fmt.Printf("%+v ", typeof3(n))
	fmt.Printf("\n")

	fmt.Printf("%+v ", typeof1(f))
	fmt.Printf("%+v ", typeof2(f))
	fmt.Printf("%+v ", typeof3(f))
	fmt.Printf("\n")
}
