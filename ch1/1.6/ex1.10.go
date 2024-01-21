// Fetchall fetches URLs in parallel and reports their times and sizes.
// for excersise 1.10 we will save the content of the website to a file
package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	file, err := os.OpenFile("websiteTemp.html", os.O_RDWR | os.O_CREATE, 0755)
	if err != nil {
		panic("Oh no file open failed!")
	}	

	for _, url := range os.Args[1:] {
		go fetch(url, file, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, file *os.File, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(file, resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  bytes: %d %s", secs, nbytes, url)
}