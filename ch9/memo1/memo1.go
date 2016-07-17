package main

import (
	"net/http"
	"io/ioutil"
	"time"
	"log"
	"fmt"
	"os"
	"sync"
)

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

var sema = make(chan struct{}, 20)

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}

	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			sema <- struct {}{}
			value, err := m.Get(url)
			<- sema
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}