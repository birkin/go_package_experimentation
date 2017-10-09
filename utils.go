package main

import "fmt"

func Some_func_c() string {
	return "hello there project-b from Utils.go"
}

func Some_func_d(an_int int) string {
	var message = fmt.Sprintf("value from Utils.go, `%v`", an_int)
	return message
}
