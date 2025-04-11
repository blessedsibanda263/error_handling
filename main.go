package main

import (
	"error_handling/examples"
	"fmt"
)

func main() {
	examples.ReturningAnError()
	fmt.Println()
	examples.ErrorVariables()
	fmt.Println()
	examples.WrappingErrors()
}
