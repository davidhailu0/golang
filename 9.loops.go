package main

import "fmt"

// There is no while loop in go
/* */

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
	//can be writen as while loop using
	x := 0
	for x < 10 {
		x++
	}
	// for INITIAL; ; AFTER { // do something forever
	// }

}
