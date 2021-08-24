package main

import (
	"fmt"
)

func tryRecover() {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred", err)
		} else {
			panic(r)
		}
	}()

	b := 0
	a := 5 / b
	fmt.Println(a)

	//panic(errors.New("this is a panic error"))
}

func main() {
	tryRecover()
}
