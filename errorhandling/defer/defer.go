package main

import (
	"bufio"
	"fmt"
	"jaluik.com/learn/functional/fib/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("error occurred")
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	writer := bufio.NewWriter(file)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			panic(err)
		}
	}(writer)
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		_, err := fmt.Fprintln(writer, f())
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	//tryDefer()
	writeFile("test.txt")
}
