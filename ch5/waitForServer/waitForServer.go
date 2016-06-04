package main

import (
	"time"
	"net/http"
	"log"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	if len(os.Args) == 2 {
		waitForServer(os.Args[1])
	} else {
		log.Println("Usage: wait [url]")
	}
}


func waitForServer(url string) (string, error) {
	timeout := 1 * time.Minute
	current := time.Now()
	for retries := 1; time.Now().Before(current.Add(timeout)); retries++ {
		resp, err := http.Get(url)
		if err == nil {
			content, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err == nil {
				return string(content), nil
			}
		}
		seconds := 1 * time.Second << uint(retries)
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(seconds)
	}
	return "", fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
