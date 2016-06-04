package main

import (
	"golang.org/x/net/html"
	"os"
	"net/http"
	"log"
)

func main() {
	if len(os.Args) != 3 {
		println("usage: element [url] [id]")
		os.Exit(0)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalf("error to open %s", os.Args[1])
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("error parsing %s", os.Args[1])
	}
	defer resp.Body.Close()

	node := elementById(doc, os.Args[2])
	if node != nil {
		log.Printf("found node %s", node.Attr)
	} else {
		log.Printf("not found")
	}

}

func elementById(doc *html.Node, id string) *html.Node {

	for _, attr := range doc.Attr {
		if attr.Key == "id" && attr.Val == id {
			return doc
		}
	}

	for node := doc.FirstChild; node != nil; node = node.NextSibling {
		if elementById(node, id) != nil {
			return node
		}
	}

	return nil
}
