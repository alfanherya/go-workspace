package main

import (
	"bufio" // -> ini merupakan library yang berisi sekumpulan perintah untuk input
	"fmt"   // -> ini merupakan library yang berisi sekumpulan perintah untuk print log atau apapun itu
	"os"
)

var inputString string

func main() {
	scanner := bufio.NewReader(os.Stdin)
	text, _ := scanner.ReadString('\n')
	fmt.Println("Hello World.")
	fmt.Println(text)
}

// to complete the challenge, you must save a line of input from stdin to a variable,
// print Hello world. on a single line, and finally print value of your
// variable on a second line.
