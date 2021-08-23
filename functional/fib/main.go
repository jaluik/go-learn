package main

import (
	"bufio"
	"fmt"
	"io"
	"jaluik.com/learn/functional/fib/fib"
	"strings"
)

type intGen func() int

func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	s := fmt.Sprintf("%d\n", next)
	if next > 10000 {
		return 0, io.EOF
	}
	return strings.NewReader(s).Read(p)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	printFileContents(f)

}
