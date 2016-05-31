package main

import (
	"os"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Xkcd struct {
	Month      string `json:"month"`
	Num        int `json:"num"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type Dictionary map[string]map[int]bool

var xkcds []Xkcd
var keywords Dictionary

func main() {
	args := os.Args
	usage := "\nUsage: xkcd \n\tdownload\tto download files\n" +
			 "\tindex\tto index the files\n" +
	         "\tsearch\t search by keyword"
	if len(args) <= 1 {
		log.Println(usage)
		os.Exit(0)
	}
	if args[1] == "download" {
		download()
	} else if args[1] == "index" {
		processFiles()
	} else if args[1] == "search" {
		if len(args) != 3 {
			log.Println("xkcd search [keyword]")
			os.Exit(0)
		}
		processFiles()
		buildIndex()
		v := keywords[args[2]]
		if v != nil {


		} else {
			log.Println("nothing found")
		}
	}
}

func buildIndex() {
	keywords = make(Dictionary)
	for _, xkcd := range xkcds {
		strs := strings.Split(xkcd.Transcript, " ")
		for _, str := range strs {
			if keywords[str] == nil {
				keywords[str] = make(map[int]bool)
			}
			keywords[str][xkcd.Num] = true
		}
	}
	log.Printf("indexed a total of %d keywords", len(keywords))
}

func download() {
	//var comics []Xkcd
	start := 1
	for {
		if start == 404 {
			start++
			continue
		}
		link := fmt.Sprintf("https://xkcd.com/%d/info.0.json", start)
		log.Println("downloading " + link + " ...")
		resp, err := http.Get(link); if err != nil {
			log.Fatalf("%s", err)
			os.Exit(1)
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("status code = %d for %s", resp.StatusCode, link)
			resp.Body.Close()
			break;
		}

		bytes, err := ioutil.ReadAll(resp.Body); if err != nil {
			log.Fatalf("error reading %s", resp.Body)
		}

		fileName := fmt.Sprintf("/tmp/xkcd/%d.json", start)
		err = ioutil.WriteFile(fileName, bytes, 0644)
		check(err)

		resp.Body.Close()
		start++
	}

	log.Printf("finished downloading %d comics for xkcd website", start - 1)
}

func processFiles() {
	files, err := ioutil.ReadDir("/tmp/xkcd")
	check(err)
	for _, file := range files {
		content, err := ioutil.ReadFile("/tmp/xkcd/" + file.Name())
		check(err)
		var xkcd Xkcd
		json.Unmarshal(content, &xkcd)
		xkcds = append(xkcds, xkcd)
	}
	log.Printf("processed %d files", len(files))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
