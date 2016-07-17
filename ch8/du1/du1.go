package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"flag"
	"container/heap"
	"sync"
)

type FileInfoHeap []FileInfo

func (h FileInfoHeap) Len() int { return len(h) }
func (h FileInfoHeap) Less(i, j int) bool  { return h[i].size > h[j].size }
func (h FileInfoHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FileInfoHeap) Push(x interface{}) {
	*h = append(*h, x.(FileInfo))
}

func (h *FileInfoHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var wg sync.WaitGroup
var sema = make(chan struct{}, 20)
var done = make(chan struct{})


func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileInfos := make(chan FileInfo)
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileInfos)
	}
	go func() {
		wg.Wait()
		close(fileInfos)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()


	h := &FileInfoHeap{}

	var nfiles, nbytes int64
	for fileInfo := range fileInfos {
		nfiles++
		nbytes += fileInfo.size
		heap.Push(h, fileInfo)
	}
	printDiskUsage(nfiles, nbytes)

	for i := 0; i < 10; i++ {
		v := heap.Pop(h).(FileInfo)
		fmt.Printf("%s %.1f MB\n", v.name, float64(v.size)/1e6)
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func dirents(dir string) [] os.FileInfo {
	sema <- struct {}{}
	defer func() {<- sema}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

type FileInfo struct {
	size int64
	name string
}

func walkDir(dir string, fileInfos chan<- FileInfo)  {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, fileInfos)
		} else {
			fileInfos <- FileInfo{entry.Size(), entry.Name()}
		}
	}
}
