package main

import (
	"sort"
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}


func main() {
	sorted := topsort(prereqs)
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topsort(m map[string][]string) []string {
	var result []string
	var visited map[string]bool = make(map[string]bool)
	var visitFunc func(items []string)

	visitFunc = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				visitFunc(m[item])
				result = append(result, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitFunc(keys)
	return result
}