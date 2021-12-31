package main

import (
	"fmt"
	"go-self-samples/14_testing/funcs"
)

func main() {

	fmt.Println(funcs.JoinWithCommas([]string{"apple", "orange", "pear"}))
	fmt.Println(funcs.JoinWithCommas([]string{"apple", "orange"}))
}
