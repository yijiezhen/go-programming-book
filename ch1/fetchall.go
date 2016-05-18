package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<- ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	filename := fmt.Sprintf("/tmp/%d",start.Unix())
	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("2 %v", err)
		return
	}
	nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("3 %v", err)
		return
	}
	resp.Body.Close()
	file.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
