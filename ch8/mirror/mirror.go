package main

import (
	"log"
	"net/http"
	"golang.org/x/net/html"
	"sync"
	"strings"
)


var workList = make(chan([]string), 20)

var seen = make(map[string]bool)

var wg sync.WaitGroup

func main() {
	var initial []string
	initial = append(initial, "https://golang.org")
	workList <- initial
	wg.Add(1)
	for work := range workList {
		go dfs(work)
	}
	wg.Wait()
}


func dfs(urls []string) {
	defer wg.Done()
	for _, url := range urls {
		if !seen[url] {
			log.Println(url)
			links, err := process(url); if err != nil {
				log.Fatalf("error extracting links: %s", err)
			}
			seen[url] = true
			wg.Add(1)
			workList <- links
		}
	}
}


func process(url string) ([]string, error) {
	if !strings.HasPrefix(url, "https://golang.org") {
		return nil, nil
	}
	resp, err := http.Get(url);

	if err != nil {
		return nil, err
	}

	node, err := html.Parse(resp.Body); if err != nil {
		resp.Body.Close()
		return nil, err
	}

	var links []string

	forEachElement(node, func(node * html.Node) {
		for _, attr := range node.Attr {
			if attr.Key == "a" {
				links = append(links, attr.Val)
				break
			}
		}
	}, nil)

	//path := resp.Request.URL.Path
	resp.Body.Close()
	return links, nil
}


func forEachElement(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for nd := node.FirstChild; nd != nil; nd = nd.NextSibling {
		forEachElement(nd, pre, post)
	}

	if post != nil {
		post(node)
	}
}



