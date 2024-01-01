// prints long argument to test runtime for excercise 3
package main

import "fmt"

func main() {
	fmt.Print("ex1.3.exe ")
	for i := 0; i < 100_000; i++ {
		fmt.Print("a ")
	}
}