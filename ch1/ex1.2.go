// ex1.2 prints index and value of each argument, separated by newlines.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}