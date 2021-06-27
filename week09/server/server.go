package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "3333", "port")

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRead(conn net.Conn, n int) ([]byte, error) {
	data := make([]byte, n)
	for x := 0; x < int(n); {
		dataLen, err := conn.Read(data[x:])
		if dataLen == 0 {
			return nil, errors.New("closed")
		}
		if err != nil {
			return nil, err
		}
		x += dataLen
	}
	return data, nil
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		//io.Copy(conn, conn)

		//Read Length
		buf, err := handleRead(conn, 4)
		if err != nil {
			//fmt.Printf("Received  Len error %x\n", err)
			return
		}

		reqLen, _ := strconv.Atoi(string(buf))

		fmt.Printf("Received message Len %d\n", reqLen)
		// Handle connections in a new goroutine.

		//Read Data
		buf, err = handleRead(conn, reqLen)
		if err != nil {
			return
		}
		fmt.Println(string(buf[:reqLen-1]))

		//Write Back
		conn.Write([]byte("0007"))
		conn.Write([]byte("hello" + "\r\n"))

	}
}
