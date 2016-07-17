package main

import (
	"sync"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
	"log"
	"os"
)

type entry struct  {
	res result
	ready chan struct{}
}

type result struct {
	value interface{}
	err error
}

type Func func(key string) (interface{}, error)

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f:f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<- e.ready
	}
	return e.res.value, e.res.err
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
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}
