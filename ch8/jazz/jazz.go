package main

import (
	"net"
	"log"
	"io"
	"time"
	"io/ioutil"
)

func main() {
	conn, err := net.Dial("tcp", "beta-tn3270.loyal3.net:23")
	time.Sleep(1 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	//go func() {
	//
	//	log.Println("done")
	//} ()
	io.WriteString(conn, "string PFLU043\r")
	bytes, err := ioutil.ReadAll(conn)
	println(string(bytes))
	io.WriteString(conn, "string tab\r")
	io.WriteString(conn, "string pass1234\r")
	io.WriteString(conn, "enter\r")
	io.WriteString(conn, "ascii\r")
	time.Sleep(1 * time.Second)
	//mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}