// ex1.3 measures runtime difference between inefficient and efficient 
// echo program examples
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// ineffcient example echo2
	startTime1 := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		// concatenating to string makes new string, concatenates stuff, and then assigns 
		// the new string to s. The old s string memory is freed by garbage collector
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	
	duration1 := time.Now().Sub(startTime1)
	fmt.Println("Runtime Method 1: ", duration1)

	// more efficient example echo3
	startTime2 := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	duration2 := time.Now().Sub(startTime2)
	fmt.Println("Runtime Method 2: ", duration2)
}