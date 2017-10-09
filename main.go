package main

import "fmt"
import "project_b_stuff/project_b/libs"

func main() {

	fmt.Println("hello project-b")

	var somestring_a string
	somestring_a = libs.Some_func_a()
	fmt.Println(fmt.Sprintf("somestring_a, ```%v```", somestring_a))

	var somestring_b string
	somestring_b = libs.Some_func_b()
	fmt.Println(fmt.Sprintf("somestring_b, ```%v```", somestring_b))

	var somestring_c string
	somestring_c = Some_func_c()
	fmt.Println(fmt.Sprintf("somestring_c, ```%v```", somestring_c))

	var somestring_d string
	var someint_d int = 42
	somestring_d = Some_func_d(someint_d)
	fmt.Println(fmt.Sprintf("somestring_d, ```%v```", somestring_d))

	var somestring_e1 string
	// var somestring_e2 string
	type SomeObject struct {
		foo_a string
	}
	some_object := SomeObject{"string-42"}
	somestring_e1 = some_object.foo_a
	fmt.Println(fmt.Sprintf("somestring_e1, ```%v```", somestring_e1))
	// somestring_e2 = Some_func_e(some_object)
	// fmt.Println(fmt.Sprintf("somestring_e2, ```%v```", somestring_e2))

}

/*
running `go run ./*.go` yields...

hello project-b
somestring_a, ```hello there project-b from file_a.go```
somestring_b, ```hello there project-b from file_b.go```
somestring_c, ```hello there project-b from Utils.go```
somestring_d, ```value from Utils.go, `42````
somestring_e1, ```string-42```
*/
