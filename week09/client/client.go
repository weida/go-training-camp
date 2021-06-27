package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3333", "port")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	done := make(chan string)
	go handleWrite(conn, done)
	go handleRead(conn, done)
	fmt.Println(<-done)
	fmt.Println(<-done)
}
func handleWrite(conn net.Conn, done chan string) {
	_, e := conn.Write([]byte("0007"))
	if e != nil {
		fmt.Println("Error to send message because of ", e.Error())
		return
	}
	_, e = conn.Write([]byte("hello " + "\r\n"))
	if e != nil {
		fmt.Println("Error to send message because of ", e.Error())
		return
	}
	done <- "Sent"
}
func handleRead(conn net.Conn, done chan string) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	fmt.Println(string(buf[:reqLen-1]))
	done <- "Read"
}
