package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func Listen(port *int) {
	address := fmt.Sprintf(":%d",*port)
	l, err := net.Listen("tcp", address)
	if nil != err {
		fmt.Println(err)
	}
	defer l.Close()
	fmt.Printf("Listening on %d", *port)
	session := 0;
	for {
		c, _ := l.Accept()
		session++
		shell(c,session)
	}
}

func shell(c net.Conn, session int){
		fmt.Println("\nAccepted connection from", c.RemoteAddr(), "on session",session)
		go io.Copy(c, os.Stdin)
		go io.Copy(os.Stdout, c)
}

func main() {
	port := flag.Int("p",1234,"port")
	flag.Parse()
	Listen(port)
}