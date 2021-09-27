package main

import (
	"fmt"
	"os"

	"github.com/JYPark09/monkey-interpreter/repl"
)

func main() {
	fmt.Println("Monkey Interpreter")

	repl.Start(os.Stdin, os.Stdout)
}
