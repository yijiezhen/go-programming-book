package main

import (
	"net"
	"io"
	"os"
	"log"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if tc, ok; conn.(*net.TCPConn);
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Print(err)
		}
		log.Println("done")
		done <- struct {} {}
	} ()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}