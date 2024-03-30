// echo prints the name of the command and its comand-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	main1()
}

func main2() {
	fmt.Println("invoking main2()")
	fmt.Println("the command: ", os.Args[0])
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("the command line args:", s)
}

func main1() {
	fmt.Println("invoking main1()")
	fmt.Println("the command: ", os.Args[0])
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("the command line args:", s)
}
