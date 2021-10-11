package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func Listen(port *int) {
	reader := bufio.NewReader(os.Stdin)
	address := fmt.Sprintf(":%d",*port)
	l, err := net.Listen("tcp", address)
	if nil != err {
		fmt.Println(err)
	}
	defer l.Close()
	fmt.Printf("Listening on %d", *port)
	for {
		c, _ := l.Accept()
		cmd,_ := reader.ReadString('\n')
		cmd = strings.TrimSuffix(cmd,"\n")
		shell(c)
	}
}

func shell(c net.Conn){
	fmt.Println("\nAccepted connection from", c.RemoteAddr())
	go io.Copy(c, os.Stdin)
}

func main() {
	port := flag.Int("p",1234,"port")
	flag.Parse()
	Listen(port)
}